package handler

import (
	"errors"
	"fmt"
	discord "github.com/chremoas/discord-gateway/proto"
	rolesrv "github.com/chremoas/role-srv/proto"
	common "github.com/chremoas/services-common/command"
	"github.com/chremoas/services-common/config"
	redis "github.com/chremoas/services-common/redis"
	"github.com/chremoas/services-common/sets"
	"github.com/fatih/structs"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type rolesHandler struct {
	Client client.Client
	Redis  *redis.Client
	Logger *zap.Logger
}

type clientList struct {
	discord discord.DiscordGatewayService
}

var clients clientList
var botRole string
var ignoredRoles []string
var roleKeys = []string{"Name", "Color", "Hoist", "Position", "Permissions", "Managed", "Mentionable"}
var roleTypes = []string{"internal", "discord"}

func NewRolesHandler(config *config.Configuration, service micro.Service, log *zap.Logger) rolesrv.RolesHandler {
	c := service.Client()

	clients = clientList{
		discord: discord.NewDiscordGatewayService(config.LookupService("gateway", "discord"), c),
	}

	botRole = config.Bot.BotRole
	ignoredRoles = config.Bot.IgnoredRoles

	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
	redisClient := redis.Init(addr, config.Redis.Password, config.Redis.Database, config.LookupService("srv", "perms"))

	_, err := redisClient.Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &rolesHandler{Redis: redisClient, Logger: log}
}

func (h *rolesHandler) GetRoleKeys(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.StringList) error {
	response.Value = roleKeys
	return nil
}

func (h *rolesHandler) GetRoleTypes(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.StringList) error {
	response.Value = roleTypes
	return nil
}

func (h *rolesHandler) AddRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.NilMessage) error {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.ShortName))
	filterA := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.FilterA))
	filterB := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.FilterB))

	// Type, Name and the filters are required so let's check for those
	if len(request.Type) == 0 {
		return errors.New("Type is required.")
	}

	if len(request.ShortName) == 0 {
		return errors.New("ShortName is required.")
	}

	if len(request.Name) == 0 {
		return errors.New("Name is required.")
	}

	if len(request.FilterA) == 0 {
		return errors.New("FilterA is required.")
	}

	if len(request.FilterB) == 0 {
		return errors.New("FilterB is required.")
	}

	if !validListItem(request.Type, roleTypes) {
		return fmt.Errorf("`%s` isn't a valid Role Type.", request.Type)
	}

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("Role `%s` already exists.", request.Name)
	}

	// Check if filter A exists
	exists, err = h.Redis.Client.Exists(filterA).Result()

	if err != nil {
		return err
	}

	if exists == 0 && request.FilterA != "wildcard" {
		return fmt.Errorf("FilterA `%s` doesn't exists.", request.FilterA)
	}

	// Check if filter B exists
	exists, err = h.Redis.Client.Exists(filterB).Result()

	if err != nil {
		return err
	}

	if exists == 0 && request.FilterB != "wildcard" {
		return fmt.Errorf("FilterB `%s` doesn't exists.", request.FilterB)
	}

	_, err = h.Redis.Client.HMSet(roleName, structs.Map(request)).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}

	return nil
}

func (h *rolesHandler) UpdateRole(ctx context.Context, request *rolesrv.UpdateInfo, response *rolesrv.NilMessage) error {
	// Does this actually work? -brian
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.Name))

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Role `%s` doesn't exists.", request.Name)
	}

	if !validListItem(request.Key, roleKeys) {
		return fmt.Errorf("`%s` isn't a valid Role Key.", request.Key)
	}

	h.Redis.Client.HSet(roleName, request.Key, request.Value)

	return nil
}

func validListItem(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (h *rolesHandler) RemoveRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.NilMessage) error {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", request.ShortName))

	exists, err := h.Redis.Client.Exists(roleName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Role `%s` doesn't exists.", request.ShortName)
	}

	_, err = h.Redis.Client.Del(roleName).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) GetRoles(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.GetRolesResponse) error {
	var sigValue, joinableValue bool
	roles, err := h.getRoles()

	if err != nil {
		return err
	}

	for role := range roles {
		roleInfo, err := h.Redis.Client.HGetAll(h.Redis.KeyName(fmt.Sprintf("role:%s", roles[role]))).Result()
		if err != nil {
			return err
		}

		if roleInfo["Sig"] == "0" {
			sigValue = false
		} else {
			sigValue = true
		}

		if roleInfo["Joinable"] == "0" {
			joinableValue = false
		} else {
			joinableValue = true
		}

		response.Roles = append(response.Roles, &rolesrv.Role{
			ShortName: roles[role],
			Name:      roleInfo["Name"],
			Sig:       sigValue,
			Joinable:  joinableValue,
		})
	}

	return nil
}

func (h *rolesHandler) getRoles() ([]string, error) {
	var roleList []string
	roles, err := h.Redis.Client.Keys(h.Redis.KeyName("role:*")).Result()

	if err != nil {
		return nil, err
	}

	for role := range roles {
		roleName := strings.Split(roles[role], ":")
		roleList = append(roleList, roleName[len(roleName)-1])
	}

	return roleList, nil
}

func (h *rolesHandler) getRole(name string) (role map[string]string, err error) {
	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", name))

	exists, err := h.Redis.Client.Exists(roleName).Result()
	if err != nil {
		return nil, err
	}

	if exists == 0 {
		return nil, fmt.Errorf("role doesn't exist: %s", name)
	}

	r, err := h.Redis.Client.HGetAll(roleName).Result()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h *rolesHandler) GetRole(ctx context.Context, request *rolesrv.Role, response *rolesrv.Role) error {
	role, err := h.getRole(request.ShortName)

	if err != nil {
		return err
	}

	color, _ := strconv.ParseInt(role["Color"], 10, 32)
	position, _ := strconv.ParseInt(role["Position"], 10, 32)
	permissions, _ := strconv.ParseInt(role["Permissions"], 10, 32)
	hoist, _ := strconv.ParseBool(role["Hoist"])
	managed, _ := strconv.ParseBool(role["Managed"])
	mentionable, _ := strconv.ParseBool(role["Mentionable"])
	sig, _ := strconv.ParseBool(role["Sig"])
	joinable, _ := strconv.ParseBool(role["Joinable"])

	response.ShortName = request.ShortName
	response.Type = role["Type"]
	response.FilterA = role["FilterA"]
	response.FilterB = role["FilterB"]
	response.Name = role["Name"]
	response.Color = int32(color)
	response.Hoist = hoist
	response.Position = int32(position)
	response.Permissions = int32(permissions)
	response.Managed = managed
	response.Mentionable = mentionable
	response.Sig = sig
	response.Joinable = joinable

	return nil
}

func (h *rolesHandler) SyncMembers(ctx context.Context, request *rolesrv.SyncRequest, response *rolesrv.NilMessage) error {
	go h.syncMembers(request.ChannelId, request.UserId, request.SendMessage)
	return nil
}

func (h *rolesHandler) sendMessage(ctx context.Context, channelId, message string, sendMessage bool) {
	if sendMessage {
		clients.discord.SendMessage(ctx, &discord.SendMessageRequest{ChannelId: channelId, Message: message})
	}
}

func (h *rolesHandler) syncMembers(channelId, userId string, sendMessage bool) error {
	ctx := context.Background()
	sugar := h.Logger.Sugar()
	sugar.Info("Starting syncMembers()")
	h.sendMessage(ctx, channelId, common.SendSuccess("Starting member sync"), sendMessage)
	var roleNameMap = make(map[string]string)
	var membershipSets = make(map[string]*sets.StringSet)

	// Discord limit is 1000, should probably make this a config option. -brian
	var numberPerPage int32 = 1000
	var memberCount = 1
	var memberId = ""

	h.sendMessage(ctx, channelId, common.SendSuccess("Getting all Discord members"), sendMessage)
	// Need to pre-populate the membership sets with all the users so we can pick up users with no roles.
	for memberCount > 0 {
		//longCtx, _ := context.WithTimeout(context.Background(), time.Second * 20)

		members, err := clients.discord.GetAllMembers(ctx, &discord.GetAllMembersRequest{NumberPerPage: numberPerPage, After: memberId})
		if err != nil {
			msg := fmt.Sprintf("syncMembers: GetAllMembers: %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		for m := range members.Members {
			userId := members.Members[m].User.Id
			if _, ok := membershipSets[userId]; !ok {
				membershipSets[userId] = sets.NewStringSet()
			}

			oldNum, _ := strconv.Atoi(members.Members[m].User.Id)
			newNum, _ := strconv.Atoi(memberId)

			if oldNum > newNum {
				memberId = members.Members[m].User.Id
			}
		}

		memberCount = len(members.Members)
	}

	h.sendMessage(ctx, channelId, common.SendSuccess("Getting all Discord roles"), sendMessage)
	// Get all the Roles from discord and create a map of their name to their Id
	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})
	if err != nil {
		msg := fmt.Sprintf("syncMembers: GetAllRoles: %s", err.Error())
		h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
		sugar.Error(msg)
		return err
	}

	for d := range discordRoles.Roles {
		roleNameMap[discordRoles.Roles[d].Name] = discordRoles.Roles[d].Id
	}

	h.sendMessage(ctx, channelId, common.SendSuccess("Getting all Chremoas roles"), sendMessage)
	// Get all the Chremoas roles and build membership Sets
	chremoasRoles, err := h.getRoles()
	if err != nil {
		msg := fmt.Sprintf("syncMembers: getRoles: %s", err.Error())
		h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
		sugar.Error(msg)
		return err
	}

	h.sendMessage(ctx, channelId, common.SendSuccess("Getting all role membership"), sendMessage)
	for r := range chremoasRoles {
		fmt.Printf("Checking role: %s\n", chremoasRoles[r])
		membership, err := h.getRoleMembership(chremoasRoles[r])
		if err != nil {
			msg := fmt.Sprintf("syncMembers: getRoleMembership: %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		roleName, err := h.getRole(chremoasRoles[r])
		if err != nil {
			msg := fmt.Sprintf("syncMembers: getRole: %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		roleId := roleNameMap[roleName["Name"]]

		for m := range membership.Set {
			sugar.Infof("Key is: %s", m)
			if len(m) != 0 {
				sugar.Infof("Set is %v", membershipSets[m])
				if membershipSets[m] == nil {
					membershipSets[m] = sets.NewStringSet()
				}
				membershipSets[m].Add(roleId)
			}
		}
	}

	h.sendMessage(ctx, channelId, common.SendSuccess("Updating Discord Roles"), sendMessage)
	// Apply the membership sets to discord overwriting anything that's there.
	for m := range membershipSets {
		clients.discord.UpdateMember(ctx, &discord.UpdateMemberRequest{
			Operation: discord.MemberUpdateOperation_ADD_OR_UPDATE_ROLES,
			UserId:    m,
			RoleIds:   membershipSets[m].ToSlice(),
		})
	}

	h.Logger.Info("Finished syncMembers()")
	h.sendMessage(ctx, channelId, common.SendSuccess("Finished syncMembers()"), sendMessage)
	return nil
}

func (h *rolesHandler) GetRoleMembership(ctx context.Context, request *rolesrv.RoleMembershipRequest, response *rolesrv.RoleMembershipResponse) error {
	members, err := h.getRoleMembership(request.Name)
	if err != nil {
		return err
	}

	for m := range members.Set {
		response.Members = append(response.Members, m)
	}

	return nil
}

func (h *rolesHandler) getRoleMembership(role string) (members *sets.StringSet, err error) {
	var filterASet = sets.NewStringSet()
	var filterBSet = sets.NewStringSet()

	roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", role))

	r, err := h.Redis.Client.HGetAll(roleName).Result()
	if err != nil {
		return filterASet, err
	}

	filterADesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", r["FilterA"]))
	filterBDesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", r["FilterB"]))

	filterAMembers := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", r["FilterA"]))
	filterBMembers := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", r["FilterB"]))

	if r["FilterB"] == "wildcard" {
		exists, err := h.Redis.Client.Exists(filterADesc).Result()
		if err != nil {
			return filterASet, err
		}

		if exists == 0 {
			return filterASet, fmt.Errorf("Filter `%s` doesn't exists.", r["FilterA"])
		}

		filterA, err := h.Redis.Client.SMembers(filterAMembers).Result()
		if err != nil {
			return filterASet, err
		}

		filterASet.FromSlice(filterA)
		return filterASet, nil
	}

	if r["FilterA"] == "wildcard" {
		exists, err := h.Redis.Client.Exists(filterBDesc).Result()
		if err != nil {
			return filterASet, err
		}

		if exists == 0 {
			return filterASet, fmt.Errorf("Filter `%s` doesn't exists.", r["FilterB"])
		}

		filterB, err := h.Redis.Client.SMembers(filterBMembers).Result()
		if err != nil {
			return filterASet, err
		}

		filterBSet.FromSlice(filterB)
		return filterBSet, nil
	}

	filterInter, err := h.Redis.Client.SInter(filterAMembers, filterBMembers).Result()
	if err != nil {
		return filterASet, err
	}

	filterASet.FromSlice(filterInter)
	return filterASet, nil
}

func (h *rolesHandler) SyncRoles(ctx context.Context, request *rolesrv.SyncRequest, response *rolesrv.NilMessage) error {
	// This was spawning a thread but I think I want to keep this one serialized so that it's guaranteed to finish
	// before member syncing starts. -brian
	h.syncRoles(request.ChannelId, request.UserId, request.SendMessage)
	return nil
}

func (h *rolesHandler) syncRoles(channelId, userId string, sendMessage bool) error {
	ctx := context.Background()
	var matchDiscordError = regexp.MustCompile(`^The role '.*' already exists$`)
	chremoasRoleSet := sets.NewStringSet()
	discordRoleSet := sets.NewStringSet()
	sugar := h.Logger.Sugar()
	var chremoasRoleData = make(map[string]map[string]string)

	h.sendMessage(ctx, channelId, common.SendSuccess("Starting Role Sync"), sendMessage)
	sugar.Info("syncRoles: Starting Role Sync")

	chremoasRoles, err := h.getRoles()
	if err != nil {
		msg := fmt.Sprintf("syncRoles: h.getRoles(): %s", err.Error())
		h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
		sugar.Error(msg)
		return err
	}

	for role := range chremoasRoles {
		roleName := h.Redis.KeyName(fmt.Sprintf("role:%s", chremoasRoles[role]))
		c, err := h.Redis.Client.HGetAll(roleName).Result()

		if err != nil {
			msg := fmt.Sprintf("syncRoles: HGetAll(): %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		chremoasRoleSet.Add(c["Name"])

		if mm, ok := chremoasRoleData[c["Name"]]; !ok {
			mm = make(map[string]string)
			chremoasRoleData[c["Name"]] = mm
		}
		chremoasRoleData[c["Name"]] = c
	}

	discordRoles, err := clients.discord.GetAllRoles(ctx, &discord.GuildObjectRequest{})
	if err != nil {
		msg := fmt.Sprintf("syncRoles: GetAllRoles: %s", err.Error())
		h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
		sugar.Error(msg)
		return err
	}

	ignoreSet := sets.NewStringSet()
	ignoreSet.Add(botRole)
	ignoreSet.Add("@everyone")
	for i := range ignoredRoles {
		ignoreSet.Add(ignoredRoles[i])
	}

	for role := range discordRoles.Roles {
		if !ignoreSet.Contains(discordRoles.Roles[role].Name) {
			discordRoleSet.Add(discordRoles.Roles[role].Name)
		}
	}

	toAdd := chremoasRoleSet.Difference(discordRoleSet)
	toDelete := discordRoleSet.Difference(chremoasRoleSet)
	toUpdate := discordRoleSet.Intersection(chremoasRoleSet)

	sugar.Infof("toAdd: %v", toAdd)
	sugar.Infof("toDelete: %v", toDelete)
	sugar.Infof("toUpdate: %v", toUpdate)

	for r := range toAdd.Set {
		_, err := clients.discord.CreateRole(ctx, &discord.CreateRoleRequest{Name: r})

		if err != nil {
			if matchDiscordError.MatchString(err.Error()) {
				// The role list was cached most likely so we'll pretend we didn't try
				// to create it just now. -brian
				sugar.Infof("syncRoles added: %s", r)
				continue
			} else {
				msg := fmt.Sprintf("syncRoles: CreateRole(): %s", err.Error())
				h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
				sugar.Error(msg)
				return err
			}
		}

		sugar.Infof("syncRoles added: %s", r)
	}

	for r := range toDelete.Set {
		_, err := clients.discord.DeleteRole(ctx, &discord.DeleteRoleRequest{Name: r})

		if err != nil {
			msg := fmt.Sprintf("syncRoles: DeleteRole(): %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		sugar.Infof("syncRoles removed: %s", r)
	}

	for r := range toUpdate.Set {
		color, _ := strconv.ParseInt(chremoasRoleData[r]["Color"], 10, 64)
		perm, _ := strconv.ParseInt(chremoasRoleData[r]["Permissions"], 10, 64)
		position, _ := strconv.ParseInt(chremoasRoleData[r]["Position"], 10, 64)
		hoist, _ := strconv.ParseBool(chremoasRoleData[r]["Hoist"])
		mention, _ := strconv.ParseBool(chremoasRoleData[r]["Mentionable"])
		managed, _ := strconv.ParseBool(chremoasRoleData[r]["Managed"])

		editRequest := &discord.EditRoleRequest{
			Name:     chremoasRoleData[r]["Name"],
			Color:    color,
			Perm:     perm,
			Position: position,
			Hoist:    hoist,
			Mention:  mention,
			Managed:  managed,
		}

		longCtx, _ := context.WithTimeout(ctx, time.Minute*5)
		_, err := clients.discord.EditRole(longCtx, editRequest)
		if err != nil {
			msg := fmt.Sprintf("syncRoles: EditRole(): %s", err.Error())
			h.sendMessage(ctx, channelId, common.SendFatal(msg), true)
			sugar.Error(msg)
			return err
		}

		sugar.Infof("syncRoles updated: %s", r)
	}

	h.sendMessage(ctx, channelId, common.SendSuccess("Finished Role Sync"), sendMessage)
	sugar.Info("syncRoles: Finished Role Sync")

	return nil
}

//
// Filter related stuff
//

func (h *rolesHandler) GetFilters(ctx context.Context, request *rolesrv.NilMessage, response *rolesrv.FilterList) error {
	filters, err := h.Redis.Client.Keys(h.Redis.KeyName("filter_description:*")).Result()

	if err != nil {
		return err
	}

	for filter := range filters {
		filterDescription, err := h.Redis.Client.Get(filters[filter]).Result()

		if err != nil {
			return err
		}

		filterName := strings.Split(filters[filter], ":")

		response.FilterList = append(response.FilterList,
			&rolesrv.Filter{Name: filterName[len(filterName)-1], Description: filterDescription})
	}

	return nil
}

func (h *rolesHandler) AddFilter(ctx context.Context, request *rolesrv.Filter, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Name))

	// Type and Name are required so let's check for those
	if len(request.Name) == 0 {
		return errors.New("Name is required.")
	}

	if len(request.Description) == 0 {
		return errors.New("Description is required.")
	}

	exists, err := h.Redis.Client.Exists(filterName).Result()

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("Filter `%s` already exists.", request.Name)
	}

	_, err = h.Redis.Client.Set(filterName, request.Description, 0).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}

	return nil
}

func (h *rolesHandler) RemoveFilter(ctx context.Context, request *rolesrv.Filter, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Name))
	filterMembers := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Name))

	exists, err := h.Redis.Client.Exists(filterName).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Name)
	}

	members, err := h.Redis.Client.SMembers(filterMembers).Result()

	if len(members) > 0 {
		return fmt.Errorf("Filter `%s` not empty.", request.Name)
	}

	_, err = h.Redis.Client.Del(filterName).Result()

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) GetMembers(ctx context.Context, request *rolesrv.Filter, response *rolesrv.MemberList) error {
	var memberlist []string
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Name))

	filters, err := h.Redis.Client.SMembers(filterName).Result()

	if err != nil {
		return err
	}

	for filter := range filters {
		if len(filters[filter]) > 0 {
			memberlist = append(memberlist, filters[filter])
		}
	}

	response.Members = memberlist
	return nil
}

func (h *rolesHandler) AddMembers(ctx context.Context, request *rolesrv.Members, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Filter))
	filterDesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Filter))

	exists, err := h.Redis.Client.Exists(filterDesc).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Filter)
	}

	for member := range request.Name {
		_, err = h.Redis.Client.SAdd(filterName, request.Name[member]).Result()
	}

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) RemoveMembers(ctx context.Context, request *rolesrv.Members, response *rolesrv.NilMessage) error {
	filterName := h.Redis.KeyName(fmt.Sprintf("filter_members:%s", request.Filter))
	filterDesc := h.Redis.KeyName(fmt.Sprintf("filter_description:%s", request.Filter))

	exists, err := h.Redis.Client.Exists(filterDesc).Result()

	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("Filter `%s` doesn't exists.", request.Filter)
	}

	for member := range request.Name {
		_, err = h.Redis.Client.SRem(filterName, request.Name[member]).Result()
	}

	if err != nil {
		return err
	}

	response = &rolesrv.NilMessage{}
	return nil
}

func (h *rolesHandler) GetDiscordUser(ctx context.Context, request *rolesrv.GetDiscordUserRequest, response *rolesrv.GetDiscordUserResponse) error {
	user, err := clients.discord.GetUser(ctx, &discord.GetUserRequest{UserId: request.UserId})
	if err != nil {
		return err
	}

	response.Username = user.User.Username
	response.Avatar = user.User.Avatar
	response.Bot = user.User.Bot
	response.Discriminator = user.User.Discriminator
	response.Email = user.User.Email
	response.MfaEnabled = user.User.MFAEnabled
	response.Verified = user.User.Verified

	return nil
}
