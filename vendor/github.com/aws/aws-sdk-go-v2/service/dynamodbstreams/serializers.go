// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson10_serializeOpDescribeStream struct {
}

func (*awsAwsjson10_serializeOpDescribeStream) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeStreamInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.DescribeStream")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeStreamInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetRecords struct {
}

func (*awsAwsjson10_serializeOpGetRecords) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRecordsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.GetRecords")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetRecordsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetShardIterator struct {
}

func (*awsAwsjson10_serializeOpGetShardIterator) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetShardIterator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetShardIteratorInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.GetShardIterator")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetShardIteratorInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListStreams struct {
}

func (*awsAwsjson10_serializeOpListStreams) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListStreams) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListStreamsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.ListStreams")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListStreamsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeOpDocumentDescribeStreamInput(v *DescribeStreamInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExclusiveStartShardId != nil {
		ok := object.Key("ExclusiveStartShardId")
		ok.String(*v.ExclusiveStartShardId)
	}

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.StreamArn != nil {
		ok := object.Key("StreamArn")
		ok.String(*v.StreamArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetRecordsInput(v *GetRecordsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.ShardIterator != nil {
		ok := object.Key("ShardIterator")
		ok.String(*v.ShardIterator)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetShardIteratorInput(v *GetShardIteratorInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SequenceNumber != nil {
		ok := object.Key("SequenceNumber")
		ok.String(*v.SequenceNumber)
	}

	if v.ShardId != nil {
		ok := object.Key("ShardId")
		ok.String(*v.ShardId)
	}

	if len(v.ShardIteratorType) > 0 {
		ok := object.Key("ShardIteratorType")
		ok.String(string(v.ShardIteratorType))
	}

	if v.StreamArn != nil {
		ok := object.Key("StreamArn")
		ok.String(*v.StreamArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListStreamsInput(v *ListStreamsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExclusiveStartStreamArn != nil {
		ok := object.Key("ExclusiveStartStreamArn")
		ok.String(*v.ExclusiveStartStreamArn)
	}

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}
