// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func connection0(s Builder) {
	s.Auth().Username("1").Password("1").Build()
	s.Auth().Password("1").Build()
	s.ClientCaching().Yes().Build()
	s.ClientCaching().No().Build()
	s.ClientGetname().Build()
	s.ClientGetredir().Build()
	s.ClientId().Build()
	s.ClientInfo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").Laddr("1").SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").Laddr("1").SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().User("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeNormal().Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().User("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeMaster().Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().User("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypeReplica().Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().User("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).TypePubsub().Build()
	s.ClientKill().IpPort("1").Id(1).User("1").Build()
	s.ClientKill().IpPort("1").Id(1).Addr("1").Build()
	s.ClientKill().IpPort("1").Id(1).Laddr("1").Build()
	s.ClientKill().IpPort("1").Id(1).SkipmeYes().Build()
	s.ClientKill().IpPort("1").Id(1).SkipmeNo().Build()
	s.ClientKill().IpPort("1").Id(1).Build()
	s.ClientKill().IpPort("1").TypeNormal().Build()
	s.ClientKill().IpPort("1").TypeMaster().Build()
	s.ClientKill().IpPort("1").TypeReplica().Build()
	s.ClientKill().IpPort("1").TypePubsub().Build()
	s.ClientKill().IpPort("1").User("1").Build()
	s.ClientKill().IpPort("1").Addr("1").Build()
	s.ClientKill().IpPort("1").Laddr("1").Build()
	s.ClientKill().IpPort("1").SkipmeYes().Build()
	s.ClientKill().IpPort("1").SkipmeNo().Build()
	s.ClientKill().IpPort("1").Build()
	s.ClientKill().Id(1).Build()
	s.ClientKill().TypeNormal().Build()
	s.ClientKill().TypeMaster().Build()
	s.ClientKill().TypeReplica().Build()
	s.ClientKill().TypePubsub().Build()
	s.ClientKill().User("1").Build()
	s.ClientKill().Addr("1").Build()
	s.ClientKill().Laddr("1").Build()
	s.ClientKill().SkipmeYes().Build()
	s.ClientKill().SkipmeNo().Build()
	s.ClientKill().Build()
	s.ClientList().TypeNormal().Id().ClientId(1).ClientId(1).Build()
	s.ClientList().TypeNormal().Build()
	s.ClientList().TypeMaster().Id().ClientId(1).ClientId(1).Build()
	s.ClientList().TypeMaster().Build()
	s.ClientList().TypeReplica().Id().ClientId(1).ClientId(1).Build()
	s.ClientList().TypeReplica().Build()
	s.ClientList().TypePubsub().Id().ClientId(1).ClientId(1).Build()
	s.ClientList().TypePubsub().Build()
	s.ClientList().Id().ClientId(1).ClientId(1).Build()
	s.ClientList().Build()
	s.ClientNoEvict().On().Build()
	s.ClientNoEvict().Off().Build()
	s.ClientNoTouch().On().Build()
	s.ClientNoTouch().Off().Build()
	s.ClientPause().Timeout(1).Write().Build()
	s.ClientPause().Timeout(1).All().Build()
	s.ClientPause().Timeout(1).Build()
	s.ClientReply().On().Build()
	s.ClientReply().Off().Build()
	s.ClientReply().Skip().Build()
	s.ClientSetinfo().Libname("1").Build()
	s.ClientSetinfo().Libver("1").Build()
	s.ClientSetname().ConnectionName("1").Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Optin().Optout().Noloop().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Optin().Optout().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Optin().Noloop().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Optin().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Optout().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Noloop().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Bcast().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Optin().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Optout().Build()
}

func connection1(s Builder) {
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Noloop().Build()
	s.ClientTracking().On().Redirect(1).Prefix().Prefix("1").Prefix("1").Build()
	s.ClientTracking().On().Redirect(1).Bcast().Build()
	s.ClientTracking().On().Redirect(1).Optin().Build()
	s.ClientTracking().On().Redirect(1).Optout().Build()
	s.ClientTracking().On().Redirect(1).Noloop().Build()
	s.ClientTracking().On().Redirect(1).Build()
	s.ClientTracking().On().Prefix().Prefix("1").Prefix("1").Build()
	s.ClientTracking().On().Bcast().Build()
	s.ClientTracking().On().Optin().Build()
	s.ClientTracking().On().Optout().Build()
	s.ClientTracking().On().Noloop().Build()
	s.ClientTracking().On().Build()
	s.ClientTracking().Off().Redirect(1).Build()
	s.ClientTracking().Off().Prefix().Prefix("1").Prefix("1").Build()
	s.ClientTracking().Off().Bcast().Build()
	s.ClientTracking().Off().Optin().Build()
	s.ClientTracking().Off().Optout().Build()
	s.ClientTracking().Off().Noloop().Build()
	s.ClientTracking().Off().Build()
	s.ClientTrackinginfo().Build()
	s.ClientUnblock().ClientId(1).Timeout().Build()
	s.ClientUnblock().ClientId(1).Error().Build()
	s.ClientUnblock().ClientId(1).Build()
	s.ClientUnpause().Build()
	s.Echo().Message("1").Build()
	s.Hello().Protover(1).Auth("1", "1").Setname("1").Build()
	s.Hello().Protover(1).Auth("1", "1").Build()
	s.Hello().Protover(1).Setname("1").Build()
	s.Hello().Protover(1).Build()
	s.Hello().Build()
	s.Ping().Message("1").Build()
	s.Ping().Build()
	s.Quit().Build()
	s.Reset().Build()
	s.Select().Index(1).Build()
}

func TestCommand_InitSlot_connection(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { connection0(s) })
	t.Run("1", func(t *testing.T) { connection1(s) })
}

func TestCommand_NoSlot_connection(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { connection0(s) })
	t.Run("1", func(t *testing.T) { connection1(s) })
}
