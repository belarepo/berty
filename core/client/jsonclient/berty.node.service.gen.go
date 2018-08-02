// this file was generated by protoc-gen-gotemplate

package jsonclient

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/berty/berty/core/api/node"
	"github.com/berty/berty/core/client"
	"github.com/berty/berty/core/entity"
)

func init() {
	registerServerStream("berty.node.EventStream", NodeEventStream)
	registerServerStream("berty.node.EventList", NodeEventList)
	registerUnary("berty.node.ContactRequest", NodeContactRequest)
	registerUnary("berty.node.ContactAcceptRequest", NodeContactAcceptRequest)
	registerUnary("berty.node.ContactRemove", NodeContactRemove)
	registerUnary("berty.node.ContactUpdate", NodeContactUpdate)
	registerServerStream("berty.node.ContactList", NodeContactList)
	registerUnary("berty.node.ConversationCreate", NodeConversationCreate)
	registerServerStream("berty.node.ConversationList", NodeConversationList)
	registerUnary("berty.node.ConversationInvite", NodeConversationInvite)
}

func NodeEventStream(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "EventStream"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().EventStream(ctx, &typedInput)
	if err != nil {
		return nil, err
	}

	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()

	return streamProxy, nil
}

func NodeEventList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "EventList"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().EventList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}

	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()

	return streamProxy, nil
}

func NodeContactRequest(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactRequest"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput node.ContactRequestInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactRequest(ctx, &typedInput)
}

func NodeContactAcceptRequest(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactAcceptRequest"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactAcceptRequest(ctx, &typedInput)
}

func NodeContactRemove(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactRemove"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactRemove(ctx, &typedInput)
}

func NodeContactUpdate(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactUpdate"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactUpdate(ctx, &typedInput)
}

func NodeContactList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactList"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().ContactList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}

	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()

	return streamProxy, nil
}

func NodeConversationCreate(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationCreate"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput entity.Conversation
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationCreate(ctx, &typedInput)
}

func NodeConversationList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationList"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().ConversationList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}

	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()

	return streamProxy, nil
}

func NodeConversationInvite(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	zap.L().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationInvite"),
		zap.String("input", string(jsonInput)),
	)

	var typedInput entity.ConversationMember
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationInvite(ctx, &typedInput)
}
