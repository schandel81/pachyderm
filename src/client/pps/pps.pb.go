// Code generated by protoc-gen-go.
// source: client/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	client/pps/pps.proto

It has these top-level messages:
	Transform
	Job
	Method
	JobInput
	JobInfo
	JobInfos
	Pipeline
	PipelineInput
	PipelineInfo
	PipelineInfos
	CreateJobRequest
	InspectJobRequest
	ListJobRequest
	GetLogsRequest
	CreatePipelineRequest
	InspectPipelineRequest
	ListPipelineRequest
	DeletePipelineRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type JobState int32

const (
	JobState_JOB_PULLING JobState = 0
	JobState_JOB_RUNNING JobState = 1
	JobState_JOB_FAILURE JobState = 2
	JobState_JOB_SUCCESS JobState = 3
)

var JobState_name = map[int32]string{
	0: "JOB_PULLING",
	1: "JOB_RUNNING",
	2: "JOB_FAILURE",
	3: "JOB_SUCCESS",
}
var JobState_value = map[string]int32{
	"JOB_PULLING": 0,
	"JOB_RUNNING": 1,
	"JOB_FAILURE": 2,
	"JOB_SUCCESS": 3,
}

func (x JobState) String() string {
	return proto.EnumName(JobState_name, int32(x))
}
func (JobState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Partition int32

const (
	Partition_BLOCK Partition = 0
	Partition_FILE  Partition = 1
	Partition_REPO  Partition = 2
)

var Partition_name = map[int32]string{
	0: "BLOCK",
	1: "FILE",
	2: "REPO",
}
var Partition_value = map[string]int32{
	"BLOCK": 0,
	"FILE":  1,
	"REPO":  2,
}

func (x Partition) String() string {
	return proto.EnumName(Partition_name, int32(x))
}
func (Partition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PipelineState int32

const (
	// When the pipeline is not ready to be triggered by commits.
	// This happens when either 1) a pipeline has been created but not yet picked
	// up by a PPS server, or 2) the pipeline does not have any inputs and is meant
	// to be triggered manually
	PipelineState_PIPELINE_IDLE PipelineState = 0
	// After this pipeline is picked up by a pachd node.  This is the normal
	// state of a pipeline.
	PipelineState_PIPELINE_RUNNING PipelineState = 1
	// After some error caused runPipeline to exit, but before the pipeline is
	// re-run.  This is when the exponential backoff is in effect.
	PipelineState_PIPELINE_RESTARTING PipelineState = 2
	// We have retried too many times and we have given up on this pipeline.
	PipelineState_PIPELINE_FAILURE PipelineState = 3
)

var PipelineState_name = map[int32]string{
	0: "PIPELINE_IDLE",
	1: "PIPELINE_RUNNING",
	2: "PIPELINE_RESTARTING",
	3: "PIPELINE_FAILURE",
}
var PipelineState_value = map[string]int32{
	"PIPELINE_IDLE":       0,
	"PIPELINE_RUNNING":    1,
	"PIPELINE_RESTARTING": 2,
	"PIPELINE_FAILURE":    3,
}

func (x PipelineState) String() string {
	return proto.EnumName(PipelineState_name, int32(x))
}
func (PipelineState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Transform struct {
	Image            string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Cmd              []string `protobuf:"bytes,2,rep,name=cmd" json:"cmd,omitempty"`
	Stdin            []string `protobuf:"bytes,3,rep,name=stdin" json:"stdin,omitempty"`
	AcceptReturnCode []int64  `protobuf:"varint,4,rep,name=accept_return_code,json=acceptReturnCode" json:"accept_return_code,omitempty"`
	Debug            bool     `protobuf:"varint,5,opt,name=debug" json:"debug,omitempty"`
}

func (m *Transform) Reset()                    { *m = Transform{} }
func (m *Transform) String() string            { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()               {}
func (*Transform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Job struct {
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Method struct {
	Partition   Partition `protobuf:"varint,1,opt,name=partition,enum=pachyderm.pps.Partition" json:"partition,omitempty"`
	Incremental bool      `protobuf:"varint,2,opt,name=incremental" json:"incremental,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type JobInput struct {
	Commit *pfs.Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Method *Method     `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
}

func (m *JobInput) Reset()                    { *m = JobInput{} }
func (m *JobInput) String() string            { return proto.CompactTextString(m) }
func (*JobInput) ProtoMessage()               {}
func (*JobInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JobInput) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *JobInput) GetMethod() *Method {
	if m != nil {
		return m.Method
	}
	return nil
}

type JobInfo struct {
	Job          *Job                        `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Transform    *Transform                  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Pipeline     *Pipeline                   `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline,omitempty"`
	Parallelism  uint64                      `protobuf:"varint,4,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs       []*JobInput                 `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob    *Job                        `protobuf:"bytes,6,opt,name=parent_job,json=parentJob" json:"parent_job,omitempty"`
	Started      *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=finished" json:"finished,omitempty"`
	OutputCommit *pfs.Commit                 `protobuf:"bytes,9,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
	State        JobState                    `protobuf:"varint,10,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobInfo) Reset()                    { *m = JobInfo{} }
func (m *JobInfo) String() string            { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()               {}
func (*JobInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *JobInfo) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *JobInfo) GetInputs() []*JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobInfo) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetStarted() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *JobInfo) GetFinished() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info,json=jobInfo" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()                    { *m = JobInfos{} }
func (m *JobInfos) String() string            { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()               {}
func (*JobInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type Pipeline struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Pipeline) Reset()                    { *m = Pipeline{} }
func (m *Pipeline) String() string            { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()               {}
func (*Pipeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PipelineInput struct {
	Repo   *pfs.Repo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Method *Method   `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
}

func (m *PipelineInput) Reset()                    { *m = PipelineInput{} }
func (m *PipelineInput) String() string            { return proto.CompactTextString(m) }
func (*PipelineInput) ProtoMessage()               {}
func (*PipelineInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PipelineInput) GetRepo() *pfs.Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *PipelineInput) GetMethod() *Method {
	if m != nil {
		return m.Method
	}
	return nil
}

type PipelineInfo struct {
	Pipeline    *Pipeline                   `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform   *Transform                  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Parallelism uint64                      `protobuf:"varint,3,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs      []*PipelineInput            `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	OutputRepo  *pfs.Repo                   `protobuf:"bytes,5,opt,name=output_repo,json=outputRepo" json:"output_repo,omitempty"`
	CreatedAt   *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	State       PipelineState               `protobuf:"varint,7,opt,name=state,enum=pachyderm.pps.PipelineState" json:"state,omitempty"`
	RecentError string                      `protobuf:"bytes,8,opt,name=recent_error,json=recentError" json:"recent_error,omitempty"`
	JobCounts   map[int32]int32             `protobuf:"bytes,9,rep,name=job_counts,json=jobCounts" json:"job_counts,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *PipelineInfo) Reset()                    { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()               {}
func (*PipelineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PipelineInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *PipelineInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputs() []*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PipelineInfo) GetJobCounts() map[int32]int32 {
	if m != nil {
		return m.JobCounts
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info,json=pipelineInfo" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()                    { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()               {}
func (*PipelineInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

type CreateJobRequest struct {
	Transform   *Transform  `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	Pipeline    *Pipeline   `protobuf:"bytes,2,opt,name=pipeline" json:"pipeline,omitempty"`
	Parallelism uint64      `protobuf:"varint,3,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs      []*JobInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob   *Job        `protobuf:"bytes,5,opt,name=parent_job,json=parentJob" json:"parent_job,omitempty"`
	Force       bool        `protobuf:"varint,6,opt,name=force" json:"force,omitempty"`
}

func (m *CreateJobRequest) Reset()                    { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()               {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateJobRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreateJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreateJobRequest) GetInputs() []*JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateJobRequest) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

type InspectJobRequest struct {
	Job        *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	BlockState bool `protobuf:"varint,2,opt,name=block_state,json=blockState" json:"block_state,omitempty"`
}

func (m *InspectJobRequest) Reset()                    { *m = InspectJobRequest{} }
func (m *InspectJobRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectJobRequest) ProtoMessage()               {}
func (*InspectJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *InspectJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type ListJobRequest struct {
	Pipeline    *Pipeline     `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	InputCommit []*pfs.Commit `protobuf:"bytes,2,rep,name=input_commit,json=inputCommit" json:"input_commit,omitempty"`
}

func (m *ListJobRequest) Reset()                    { *m = ListJobRequest{} }
func (m *ListJobRequest) String() string            { return proto.CompactTextString(m) }
func (*ListJobRequest) ProtoMessage()               {}
func (*ListJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ListJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ListJobRequest) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

type GetLogsRequest struct {
	Job *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *GetLogsRequest) Reset()                    { *m = GetLogsRequest{} }
func (m *GetLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogsRequest) ProtoMessage()               {}
func (*GetLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetLogsRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type CreatePipelineRequest struct {
	Pipeline    *Pipeline        `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform   *Transform       `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Parallelism uint64           `protobuf:"varint,3,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs      []*PipelineInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
}

func (m *CreatePipelineRequest) Reset()                    { *m = CreatePipelineRequest{} }
func (m *CreatePipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePipelineRequest) ProtoMessage()               {}
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *CreatePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreatePipelineRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreatePipelineRequest) GetInputs() []*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type InspectPipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *InspectPipelineRequest) Reset()                    { *m = InspectPipelineRequest{} }
func (m *InspectPipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectPipelineRequest) ProtoMessage()               {}
func (*InspectPipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *InspectPipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type ListPipelineRequest struct {
}

func (m *ListPipelineRequest) Reset()                    { *m = ListPipelineRequest{} }
func (m *ListPipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPipelineRequest) ProtoMessage()               {}
func (*ListPipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type DeletePipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *DeletePipelineRequest) Reset()                    { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()               {}
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *DeletePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func init() {
	proto.RegisterType((*Transform)(nil), "pachyderm.pps.Transform")
	proto.RegisterType((*Job)(nil), "pachyderm.pps.Job")
	proto.RegisterType((*Method)(nil), "pachyderm.pps.Method")
	proto.RegisterType((*JobInput)(nil), "pachyderm.pps.JobInput")
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.JobInfos")
	proto.RegisterType((*Pipeline)(nil), "pachyderm.pps.Pipeline")
	proto.RegisterType((*PipelineInput)(nil), "pachyderm.pps.PipelineInput")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.PipelineInfos")
	proto.RegisterType((*CreateJobRequest)(nil), "pachyderm.pps.CreateJobRequest")
	proto.RegisterType((*InspectJobRequest)(nil), "pachyderm.pps.InspectJobRequest")
	proto.RegisterType((*ListJobRequest)(nil), "pachyderm.pps.ListJobRequest")
	proto.RegisterType((*GetLogsRequest)(nil), "pachyderm.pps.GetLogsRequest")
	proto.RegisterType((*CreatePipelineRequest)(nil), "pachyderm.pps.CreatePipelineRequest")
	proto.RegisterType((*InspectPipelineRequest)(nil), "pachyderm.pps.InspectPipelineRequest")
	proto.RegisterType((*ListPipelineRequest)(nil), "pachyderm.pps.ListPipelineRequest")
	proto.RegisterType((*DeletePipelineRequest)(nil), "pachyderm.pps.DeletePipelineRequest")
	proto.RegisterEnum("pachyderm.pps.JobState", JobState_name, JobState_value)
	proto.RegisterEnum("pachyderm.pps.Partition", Partition_name, Partition_value)
	proto.RegisterEnum("pachyderm.pps.PipelineState", PipelineState_name, PipelineState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for API service

type APIClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (API_GetLogsClient, error)
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// DeleteAll deletes everything
	DeleteAll(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/ListJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (API_GetLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/pachyderm.pps.API/GetLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_GetLogsClient interface {
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ClientStream
}

type aPIGetLogsClient struct {
	grpc.ClientStream
}

func (x *aPIGetLogsClient) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/InspectPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/ListPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/DeletePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteAll(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/DeleteAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	InspectJob(context.Context, *InspectJobRequest) (*JobInfo, error)
	ListJob(context.Context, *ListJobRequest) (*JobInfos, error)
	GetLogs(*GetLogsRequest, API_GetLogsServer) error
	CreatePipeline(context.Context, *CreatePipelineRequest) (*google_protobuf.Empty, error)
	InspectPipeline(context.Context, *InspectPipelineRequest) (*PipelineInfo, error)
	ListPipeline(context.Context, *ListPipelineRequest) (*PipelineInfos, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*google_protobuf.Empty, error)
	// DeleteAll deletes everything
	DeleteAll(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).InspectJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/InspectJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).InspectJob(ctx, req.(*InspectJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/ListJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListJob(ctx, req.(*ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).GetLogs(m, &aPIGetLogsServer{stream})
}

type API_GetLogsServer interface {
	Send(*google_protobuf2.BytesValue) error
	grpc.ServerStream
}

type aPIGetLogsServer struct {
	grpc.ServerStream
}

func (x *aPIGetLogsServer) Send(m *google_protobuf2.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _API_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_InspectPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).InspectPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/InspectPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).InspectPipeline(ctx, req.(*InspectPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/ListPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListPipeline(ctx, req.(*ListPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/DeletePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.API/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteAll(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _API_CreateJob_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _API_InspectJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _API_ListJob_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _API_CreatePipeline_Handler,
		},
		{
			MethodName: "InspectPipeline",
			Handler:    _API_InspectPipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _API_ListPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _API_DeletePipeline_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _API_DeleteAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _API_GetLogs_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 1240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x57, 0x6b, 0x6e, 0xdb, 0xc6,
	0x13, 0xb7, 0x44, 0xbd, 0x38, 0x92, 0x15, 0x65, 0x13, 0x27, 0x82, 0x9c, 0xc4, 0xfe, 0xf3, 0x9f,
	0x02, 0x81, 0xd0, 0xca, 0xa9, 0x13, 0x04, 0x6d, 0xd1, 0x00, 0xb5, 0x15, 0x25, 0x55, 0xaa, 0xd8,
	0xea, 0xda, 0x6e, 0x81, 0x00, 0xad, 0x40, 0x51, 0x2b, 0x99, 0x0e, 0x5f, 0x25, 0x57, 0x2d, 0x7c,
	0x88, 0x5e, 0xa0, 0x3d, 0x46, 0xaf, 0xd3, 0x4b, 0xf4, 0x06, 0x9d, 0x5d, 0x92, 0x7a, 0x50, 0x8f,
	0x3a, 0x6e, 0x3e, 0xf4, 0x83, 0x80, 0xdd, 0x99, 0xd9, 0xd9, 0x99, 0xdf, 0xfc, 0x66, 0x96, 0x82,
	0xdb, 0x86, 0x65, 0x32, 0x87, 0xef, 0x79, 0x5e, 0x20, 0x7e, 0x0d, 0xcf, 0x77, 0xb9, 0x4b, 0x36,
	0x3d, 0xdd, 0x38, 0xbf, 0x1c, 0x30, 0xdf, 0x6e, 0xa0, 0xb0, 0xb6, 0x3d, 0x72, 0xdd, 0x91, 0xc5,
	0xf6, 0xa4, 0xb2, 0x3f, 0x1e, 0xee, 0x31, 0xdb, 0xe3, 0x97, 0xa1, 0x6d, 0x6d, 0x27, 0xa9, 0xe4,
	0xa6, 0xcd, 0x02, 0xae, 0xdb, 0x5e, 0x64, 0xf0, 0x20, 0x69, 0xf0, 0x8b, 0xaf, 0x7b, 0x1e, 0xf3,
	0xa3, 0xcb, 0x6a, 0x93, 0x10, 0x86, 0x81, 0xf8, 0x85, 0x52, 0xed, 0xd7, 0x14, 0xa8, 0xa7, 0xbe,
	0xee, 0x04, 0x43, 0xd7, 0xb7, 0xc9, 0x6d, 0xc8, 0x9a, 0xb6, 0x3e, 0x62, 0xd5, 0xd4, 0x6e, 0xea,
	0x91, 0x4a, 0xc3, 0x0d, 0xa9, 0x80, 0x62, 0xd8, 0x83, 0x6a, 0x7a, 0x57, 0x41, 0x99, 0x58, 0x0a,
	0xbb, 0x80, 0x0f, 0x4c, 0xa7, 0xaa, 0x48, 0x59, 0xb8, 0x21, 0x1f, 0x03, 0xd1, 0x0d, 0x83, 0x79,
	0xbc, 0xe7, 0x33, 0x3e, 0xf6, 0x9d, 0x9e, 0xe1, 0x0e, 0x58, 0x35, 0x83, 0x26, 0x0a, 0xad, 0x84,
	0x1a, 0x2a, 0x15, 0x4d, 0x94, 0x0b, 0x1f, 0x03, 0xd6, 0x1f, 0x8f, 0xaa, 0x59, 0xbc, 0xab, 0x40,
	0xc3, 0x8d, 0xb6, 0x05, 0xca, 0x6b, 0xb7, 0x4f, 0xca, 0x90, 0x36, 0x07, 0x51, 0x14, 0xb8, 0xd2,
	0xfa, 0x90, 0x7b, 0xc3, 0xf8, 0xb9, 0x3b, 0x20, 0xcf, 0x40, 0xf5, 0x74, 0x9f, 0x9b, 0xdc, 0x74,
	0x1d, 0x69, 0x50, 0xde, 0xaf, 0x36, 0xe6, 0x70, 0x6c, 0x74, 0x63, 0x3d, 0x9d, 0x9a, 0x92, 0x5d,
	0x28, 0x9a, 0x8e, 0xe1, 0x33, 0x1b, 0x31, 0xd0, 0x2d, 0x4c, 0x46, 0x5c, 0x3a, 0x2b, 0xd2, 0x7e,
	0x84, 0x02, 0x5e, 0xdd, 0x76, 0xbc, 0x31, 0x27, 0xff, 0x87, 0x9c, 0xe1, 0xda, 0xb6, 0xc9, 0xe5,
	0x15, 0xc5, 0xfd, 0x62, 0x43, 0x40, 0xd6, 0x94, 0x22, 0x1a, 0xa9, 0xc8, 0x27, 0x90, 0xb3, 0x65,
	0x50, 0xd2, 0x5b, 0x71, 0x7f, 0x2b, 0x11, 0x47, 0x18, 0x31, 0x8d, 0x8c, 0xb4, 0xbf, 0x14, 0xc8,
	0xcb, 0x0b, 0x86, 0x2e, 0x79, 0x08, 0xca, 0x85, 0xdb, 0x8f, 0x9c, 0x93, 0xc4, 0x39, 0x34, 0xa2,
	0x42, 0x2d, 0x72, 0xe5, 0x71, 0x6d, 0xa2, 0x3b, 0x92, 0xb9, 0x4e, 0x6a, 0x47, 0xa7, 0xa6, 0xe4,
	0x09, 0x14, 0x3c, 0xd3, 0x63, 0x96, 0xe9, 0x30, 0xac, 0x90, 0x38, 0x76, 0x37, 0x09, 0x51, 0xa4,
	0xa6, 0x13, 0x43, 0x01, 0x10, 0xa2, 0xa5, 0x5b, 0x16, 0x6e, 0x03, 0x1b, 0xcb, 0x96, 0x7a, 0x94,
	0xa1, 0xb3, 0x22, 0xb2, 0x07, 0x39, 0x53, 0xa0, 0x13, 0x60, 0xc9, 0x94, 0x25, 0x4e, 0x63, 0xf4,
	0x68, 0x64, 0x46, 0x3e, 0x05, 0xc0, 0xf3, 0x88, 0x6e, 0x4f, 0x24, 0x9b, 0x5b, 0x99, 0xac, 0x1a,
	0x5a, 0x89, 0xc2, 0x3f, 0x85, 0x3c, 0x92, 0xda, 0xe7, 0x6c, 0x50, 0xcd, 0x4b, 0xfb, 0x5a, 0x23,
	0xe4, 0x75, 0x23, 0xe6, 0x75, 0xe3, 0x34, 0x26, 0x3e, 0x8d, 0x4d, 0x11, 0xa8, 0xc2, 0xd0, 0x74,
	0xcc, 0xe0, 0x1c, 0x8f, 0x15, 0xfe, 0xf1, 0xd8, 0xc4, 0x96, 0x3c, 0x86, 0x4d, 0x77, 0xcc, 0x31,
	0xd6, 0x5e, 0x54, 0x6d, 0x75, 0xb1, 0xda, 0xa5, 0xd0, 0xa2, 0x19, 0xd7, 0x1c, 0xc9, 0xae, 0x73,
	0x56, 0x05, 0x49, 0xbd, 0x25, 0x10, 0x9c, 0x08, 0x35, 0x0d, 0xad, 0xb4, 0xe7, 0x11, 0xa7, 0x86,
	0xae, 0x40, 0xa3, 0x80, 0x30, 0xf4, 0x4c, 0xdc, 0x60, 0xe1, 0x05, 0x80, 0x77, 0x96, 0x01, 0x38,
	0x74, 0x69, 0xfe, 0x22, 0x5c, 0x68, 0x0f, 0xa0, 0x10, 0x57, 0x8a, 0x10, 0xc8, 0x38, 0xba, 0x1d,
	0xb7, 0xa6, 0x5c, 0x6b, 0x3f, 0xc0, 0x66, 0xac, 0x0f, 0x79, 0x7b, 0x1f, 0x32, 0x3e, 0xf3, 0xdc,
	0x88, 0x58, 0xaa, 0xcc, 0x83, 0xa2, 0x80, 0x4a, 0xf1, 0xfb, 0x32, 0xf6, 0xb7, 0x0c, 0x94, 0xa6,
	0xfe, 0x91, 0xb6, 0xb3, 0xc4, 0x4a, 0x5d, 0x95, 0x58, 0xd7, 0x65, 0x71, 0x82, 0x90, 0xca, 0x22,
	0x21, 0x9f, 0x4e, 0x08, 0x99, 0x91, 0x78, 0xde, 0x5b, 0x11, 0xcc, 0x3c, 0x2b, 0xeb, 0x50, 0x8c,
	0x8a, 0x2e, 0xa1, 0xca, 0x26, 0xa1, 0x82, 0x50, 0x2b, 0xd6, 0xe4, 0x73, 0x00, 0x1c, 0x10, 0x58,
	0xc9, 0x41, 0x4f, 0xe7, 0x11, 0x83, 0xd7, 0x51, 0x4b, 0x8d, 0xac, 0x0f, 0x38, 0xd9, 0x8f, 0x99,
	0x92, 0x97, 0x4c, 0x59, 0x15, 0xdb, 0x2c, 0x5d, 0xc8, 0xff, 0xa0, 0xe4, 0x33, 0x43, 0x34, 0x0c,
	0xf3, 0x7d, 0xd7, 0x97, 0x5c, 0x56, 0x69, 0x31, 0x94, 0xb5, 0x84, 0x88, 0xb4, 0x01, 0x04, 0x8b,
	0x0c, 0x77, 0xec, 0x60, 0xde, 0xaa, 0xcc, 0xbb, 0xbe, 0x32, 0xef, 0xa1, 0x2b, 0x48, 0xd5, 0x94,
	0xc6, 0x2d, 0x87, 0xfb, 0x97, 0x54, 0xbd, 0x88, 0xf7, 0xb5, 0x2f, 0xa1, 0x3c, 0xaf, 0x14, 0x93,
	0xfe, 0x1d, 0xbb, 0x94, 0xa5, 0xcd, 0x52, 0xb1, 0x14, 0x53, 0xfa, 0x67, 0xdd, 0x1a, 0x33, 0x59,
	0xb8, 0x2c, 0x0d, 0x37, 0x5f, 0xa4, 0x3f, 0x4b, 0x69, 0xdf, 0xce, 0x72, 0x4f, 0xf0, 0xfb, 0x2b,
	0xd8, 0x8c, 0x6b, 0x3e, 0x4b, 0xf2, 0xed, 0x35, 0xc1, 0xd1, 0x92, 0x37, 0xb3, 0xd3, 0x7e, 0x4f,
	0x43, 0xa5, 0x29, 0x01, 0x14, 0x53, 0x81, 0xfd, 0x34, 0x46, 0x54, 0xe7, 0xe9, 0x93, 0xba, 0xde,
	0x10, 0x4c, 0x5f, 0x73, 0x08, 0x2a, 0xeb, 0x86, 0x60, 0xe6, 0x3a, 0x43, 0x30, 0x7b, 0x95, 0x21,
	0x88, 0xa0, 0x63, 0x0a, 0x06, 0x93, 0x84, 0xc3, 0xa7, 0x51, 0x6e, 0xb4, 0xb7, 0x70, 0xb3, 0xed,
	0x04, 0x1e, 0x33, 0xf8, 0x0c, 0x3a, 0x57, 0x7b, 0x48, 0x76, 0xa0, 0xd8, 0xb7, 0x5c, 0xe3, 0x5d,
	0x2f, 0x64, 0x64, 0xf8, 0xf8, 0x81, 0x14, 0x49, 0xfe, 0x69, 0x63, 0x28, 0x77, 0xcc, 0x60, 0xd6,
	0xf1, 0xb5, 0x5a, 0xbd, 0x01, 0x25, 0x99, 0x75, 0x3c, 0x4e, 0xd3, 0x12, 0xa2, 0xb9, 0x71, 0x5a,
	0x94, 0x06, 0xe1, 0x46, 0x7b, 0x06, 0xe5, 0x57, 0x8c, 0x77, 0xdc, 0x51, 0xf0, 0x5e, 0xf9, 0x68,
	0x7f, 0xa6, 0x60, 0x2b, 0x24, 0xca, 0x24, 0x88, 0x7f, 0x13, 0xf6, 0x7f, 0x6c, 0x42, 0x69, 0x6f,
	0xe0, 0x4e, 0x54, 0xe9, 0x0f, 0x91, 0x1e, 0x7e, 0x53, 0xdd, 0x12, 0xc5, 0x4d, 0xf8, 0xd2, 0x3a,
	0xb0, 0xf5, 0x82, 0x59, 0xec, 0xc3, 0x60, 0x58, 0x3f, 0x92, 0x2f, 0x9d, 0x64, 0x13, 0xb9, 0x01,
	0xc5, 0xd7, 0xc7, 0x87, 0xbd, 0xee, 0x59, 0xa7, 0xd3, 0x3e, 0x7a, 0x55, 0xd9, 0x88, 0x05, 0xf4,
	0xec, 0xe8, 0x48, 0x08, 0x52, 0xb1, 0xe0, 0xe5, 0x41, 0xbb, 0x73, 0x46, 0x5b, 0x95, 0x74, 0x2c,
	0x38, 0x39, 0x6b, 0x36, 0x5b, 0x27, 0x27, 0x15, 0xa5, 0x5e, 0x07, 0x75, 0xf2, 0x1d, 0x47, 0x54,
	0xc8, 0x1e, 0x76, 0x8e, 0x9b, 0xdf, 0xa0, 0xab, 0x02, 0x64, 0x5e, 0xb6, 0x3b, 0x2d, 0xf4, 0x81,
	0x2b, 0xda, 0xea, 0x1e, 0x57, 0xd2, 0xf5, 0xd1, 0x74, 0x14, 0x85, 0x01, 0xdc, 0x44, 0x41, 0xbb,
	0xdb, 0xc2, 0xdb, 0x5b, 0xbd, 0xf6, 0x0b, 0xb4, 0xde, 0xc0, 0x9e, 0xaa, 0x4c, 0x44, 0xd3, 0x38,
	0xee, 0xc2, 0xad, 0xa9, 0xb4, 0x75, 0x72, 0x7a, 0x40, 0x4f, 0x85, 0x22, 0x3d, 0x67, 0x1e, 0x47,
	0xa9, 0xec, 0xff, 0x91, 0x05, 0xe5, 0xa0, 0xdb, 0x26, 0x87, 0xa0, 0x4e, 0xe6, 0x14, 0xd9, 0x49,
	0x80, 0x93, 0x9c, 0x60, 0xb5, 0x25, 0x34, 0xd6, 0x36, 0xc8, 0xd7, 0x00, 0xd3, 0x76, 0x26, 0xbb,
	0x09, 0x9b, 0x85, 0x4e, 0xaf, 0xad, 0xf8, 0x58, 0x40, 0x4f, 0x4d, 0xc8, 0x47, 0xcd, 0x4b, 0xee,
	0x27, 0x8c, 0xe6, 0x9b, 0xba, 0x76, 0x77, 0xb9, 0x8f, 0x00, 0x9d, 0xb4, 0x21, 0x1f, 0xb5, 0xe2,
	0x82, 0x93, 0xf9, 0x16, 0xad, 0x6d, 0x2f, 0xbc, 0x7f, 0x87, 0x97, 0x9c, 0x05, 0xdf, 0x89, 0x97,
	0x41, 0xdb, 0x78, 0x9c, 0x22, 0x5d, 0x28, 0xcf, 0x37, 0x27, 0x79, 0xb8, 0x14, 0xa2, 0x04, 0xef,
	0x30, 0xc3, 0xa4, 0xe3, 0x96, 0xf8, 0x03, 0x84, 0xc1, 0x7d, 0x0f, 0x37, 0x12, 0x0d, 0x41, 0x3e,
	0x5a, 0x0e, 0x58, 0xd2, 0xe7, 0xba, 0xd7, 0x07, 0x1d, 0x53, 0x28, 0xcd, 0xb6, 0x06, 0xd1, 0x96,
	0xe0, 0x97, 0x74, 0x79, 0x6f, 0x8d, 0x4b, 0x81, 0x24, 0xa6, 0x3f, 0xdf, 0x57, 0x0b, 0xe9, 0x2f,
	0x6d, 0xbb, 0x35, 0xe9, 0x3f, 0x07, 0x35, 0x3c, 0x72, 0x60, 0x59, 0x64, 0x85, 0xd9, 0xea, 0xe3,
	0x87, 0xd9, 0xb7, 0x0a, 0xde, 0xda, 0xcf, 0x49, 0xc5, 0x93, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0x35, 0xb7, 0xdf, 0x8c, 0x0e, 0x00, 0x00,
}
