// Code generated by svcdec; DO NOT EDIT

package crimson

import (
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

type DecoratedCrimson struct {
	// Service is the service to decorate.
	Service CrimsonServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedCrimson) ListDatacenters(c context.Context, req *ListDatacentersRequest) (rsp *ListDatacentersResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListDatacenters", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListDatacenters(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListDatacenters", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListFreeIPs(c context.Context, req *ListFreeIPsRequest) (rsp *ListIPsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListFreeIPs", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListFreeIPs(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListFreeIPs", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListOSes(c context.Context, req *ListOSesRequest) (rsp *ListOSesResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListOSes", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListOSes(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListOSes", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListPlatforms(c context.Context, req *ListPlatformsRequest) (rsp *ListPlatformsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListPlatforms", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListPlatforms(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListPlatforms", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListRacks(c context.Context, req *ListRacksRequest) (rsp *ListRacksResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListRacks", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListRacks(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListRacks", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListSwitches(c context.Context, req *ListSwitchesRequest) (rsp *ListSwitchesResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListSwitches", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListSwitches(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListSwitches", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListVLANs(c context.Context, req *ListVLANsRequest) (rsp *ListVLANsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListVLANs", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListVLANs(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListVLANs", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) CreateMachine(c context.Context, req *CreateMachineRequest) (rsp *Machine, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "CreateMachine", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.CreateMachine(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateMachine", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) DeleteMachine(c context.Context, req *DeleteMachineRequest) (rsp *google_protobuf.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "DeleteMachine", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.DeleteMachine(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "DeleteMachine", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListMachines(c context.Context, req *ListMachinesRequest) (rsp *ListMachinesResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListMachines", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListMachines(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListMachines", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) RenameMachine(c context.Context, req *RenameMachineRequest) (rsp *Machine, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "RenameMachine", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.RenameMachine(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "RenameMachine", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) UpdateMachine(c context.Context, req *UpdateMachineRequest) (rsp *Machine, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "UpdateMachine", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.UpdateMachine(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateMachine", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) CreateNIC(c context.Context, req *CreateNICRequest) (rsp *NIC, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "CreateNIC", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.CreateNIC(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateNIC", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) DeleteNIC(c context.Context, req *DeleteNICRequest) (rsp *google_protobuf.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "DeleteNIC", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.DeleteNIC(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "DeleteNIC", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListNICs(c context.Context, req *ListNICsRequest) (rsp *ListNICsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListNICs", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListNICs(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListNICs", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) UpdateNIC(c context.Context, req *UpdateNICRequest) (rsp *NIC, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "UpdateNIC", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.UpdateNIC(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateNIC", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) CreateDRAC(c context.Context, req *CreateDRACRequest) (rsp *DRAC, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "CreateDRAC", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.CreateDRAC(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateDRAC", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListDRACs(c context.Context, req *ListDRACsRequest) (rsp *ListDRACsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListDRACs", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListDRACs(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListDRACs", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) CreatePhysicalHost(c context.Context, req *CreatePhysicalHostRequest) (rsp *PhysicalHost, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "CreatePhysicalHost", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.CreatePhysicalHost(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreatePhysicalHost", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListPhysicalHosts(c context.Context, req *ListPhysicalHostsRequest) (rsp *ListPhysicalHostsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListPhysicalHosts", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListPhysicalHosts(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListPhysicalHosts", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) UpdatePhysicalHost(c context.Context, req *UpdatePhysicalHostRequest) (rsp *PhysicalHost, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "UpdatePhysicalHost", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.UpdatePhysicalHost(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdatePhysicalHost", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) CreateVM(c context.Context, req *CreateVMRequest) (rsp *VM, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "CreateVM", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.CreateVM(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "CreateVM", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) ListVMs(c context.Context, req *ListVMsRequest) (rsp *ListVMsResponse, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "ListVMs", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.ListVMs(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListVMs", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) UpdateVM(c context.Context, req *UpdateVMRequest) (rsp *VM, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "UpdateVM", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.UpdateVM(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateVM", rsp, err)
	}
	return
}

func (s *DecoratedCrimson) DeleteHost(c context.Context, req *DeleteHostRequest) (rsp *google_protobuf.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "DeleteHost", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.DeleteHost(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "DeleteHost", rsp, err)
	}
	return
}
