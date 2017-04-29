// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for Route resource provider */

// RouteToken is the type token corresponding to the Route package type.
const RouteToken = tokens.Type("aws:ec2/route:Route")

// RouteProviderOps is a pluggable interface for Route-related management functionality.
type RouteProviderOps interface {
    Check(ctx context.Context, obj *Route) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Route) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Route, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Route, new *Route, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Route, new *Route, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// RouteProvider is a dynamic gRPC-based plugin for managing Route resources.
type RouteProvider struct {
    ops RouteProviderOps
}

// NewRouteProvider allocates a resource provider that delegates to a ops instance.
func NewRouteProvider(ops RouteProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &RouteProvider{ops: ops}
}

func (p *RouteProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *RouteProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *RouteProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   string(id),
    }, nil
}

func (p *RouteProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *RouteProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    if diff.Changed("destinationCidrBlock") {
        replaces = append(replaces, "destinationCidrBlock")
    }
    if diff.Changed("routeTable") {
        replaces = append(replaces, "routeTable")
    }
    if diff.Changed("internetGateway") {
        replaces = append(replaces, "internetGateway")
    }
    if diff.Changed("vpcGatewayAttachment") {
        replaces = append(replaces, "vpcGatewayAttachment")
    }
    id := resource.ID(req.GetId())
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *RouteProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *RouteProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RouteToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *RouteProvider) Unmarshal(
    v *pbstruct.Struct) (*Route, resource.PropertyMap, mapper.DecodeError) {
    var obj Route
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Route structure(s) */

// Route is a marshalable representation of its corresponding IDL type.
type Route struct {
    Name string `json:"name"`
    DestinationCidrBlock string `json:"destinationCidrBlock"`
    RouteTable *resource.ID `json:"routeTable"`
    InternetGateway *resource.ID `json:"internetGateway"`
    VPCGatewayAttachment *resource.ID `json:"vpcGatewayAttachment"`
}

// Route's properties have constants to make dealing with diffs and property bags easier.
const (
    Route_Name = "name"
    Route_DestinationCidrBlock = "destinationCidrBlock"
    Route_RouteTable = "routeTable"
    Route_InternetGateway = "internetGateway"
    Route_VPCGatewayAttachment = "vpcGatewayAttachment"
)


