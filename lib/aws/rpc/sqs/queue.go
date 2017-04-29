// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

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

/* RPC stubs for Queue resource provider */

// QueueToken is the type token corresponding to the Queue package type.
const QueueToken = tokens.Type("aws:sqs/queue:Queue")

// QueueProviderOps is a pluggable interface for Queue-related management functionality.
type QueueProviderOps interface {
    Check(ctx context.Context, obj *Queue) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Queue) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Queue, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Queue, new *Queue, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Queue, new *Queue, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// QueueProvider is a dynamic gRPC-based plugin for managing Queue resources.
type QueueProvider struct {
    ops QueueProviderOps
}

// NewQueueProvider allocates a resource provider that delegates to a ops instance.
func NewQueueProvider(ops QueueProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &QueueProvider{ops: ops}
}

func (p *QueueProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(QueueToken))
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

func (p *QueueProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(QueueToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *QueueProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(QueueToken))
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

func (p *QueueProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(QueueToken))
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

func (p *QueueProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(QueueToken))
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
    if diff.Changed("fifoQueue") {
        replaces = append(replaces, "fifoQueue")
    }
    if diff.Changed("queueName") {
        replaces = append(replaces, "queueName")
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

func (p *QueueProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(QueueToken))
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

func (p *QueueProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(QueueToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *QueueProvider) Unmarshal(
    v *pbstruct.Struct) (*Queue, resource.PropertyMap, mapper.DecodeError) {
    var obj Queue
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Queue structure(s) */

// Queue is a marshalable representation of its corresponding IDL type.
type Queue struct {
    Name string `json:"name"`
    FIFOQueue *bool `json:"fifoQueue,omitempty"`
    QueueName *string `json:"queueName,omitempty"`
    ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty"`
    DelaySeconds *float64 `json:"delaySeconds,omitempty"`
    MaximumMessageSize *float64 `json:"maximumMessageSize,omitempty"`
    MessageRetentionPeriod *float64 `json:"messageRetentionPeriod,omitempty"`
    ReceiveMessageWaitTimeSeconds *float64 `json:"receiveMessageWaitTimeSeconds,omitempty"`
    RedrivePolicy *RedrivePolicy `json:"redrivePolicy,omitempty"`
    VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty"`
}

// Queue's properties have constants to make dealing with diffs and property bags easier.
const (
    Queue_Name = "name"
    Queue_FIFOQueue = "fifoQueue"
    Queue_QueueName = "queueName"
    Queue_ContentBasedDeduplication = "contentBasedDeduplication"
    Queue_DelaySeconds = "delaySeconds"
    Queue_MaximumMessageSize = "maximumMessageSize"
    Queue_MessageRetentionPeriod = "messageRetentionPeriod"
    Queue_ReceiveMessageWaitTimeSeconds = "receiveMessageWaitTimeSeconds"
    Queue_RedrivePolicy = "redrivePolicy"
    Queue_VisibilityTimeout = "visibilityTimeout"
)

/* Marshalable RedrivePolicy structure(s) */

// RedrivePolicy is a marshalable representation of its corresponding IDL type.
type RedrivePolicy struct {
    DeadLetterTarget *resource.ID `json:"deadLetterTarget"`
    MaxReceiveCount float64 `json:"maxReceiveCount"`
}

// RedrivePolicy's properties have constants to make dealing with diffs and property bags easier.
const (
    RedrivePolicy_DeadLetterTarget = "deadLetterTarget"
    RedrivePolicy_MaxReceiveCount = "maxReceiveCount"
)


