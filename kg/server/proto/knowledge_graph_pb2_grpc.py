# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import proto.knowledge_graph_pb2 as knowledge__graph__pb2


class KGServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetGraph = channel.unary_unary(
                '/kg.KGService/GetGraph',
                request_serializer=knowledge__graph__pb2.GetGraphReq.SerializeToString,
                response_deserializer=knowledge__graph__pb2.GetGraphRes.FromString,
                )
        self.GetItem = channel.unary_unary(
                '/kg.KGService/GetItem',
                request_serializer=knowledge__graph__pb2.GetItemReq.SerializeToString,
                response_deserializer=knowledge__graph__pb2.GetItemRes.FromString,
                )


class KGServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetGraph(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetItem(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_KGServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetGraph': grpc.unary_unary_rpc_method_handler(
                    servicer.GetGraph,
                    request_deserializer=knowledge__graph__pb2.GetGraphReq.FromString,
                    response_serializer=knowledge__graph__pb2.GetGraphRes.SerializeToString,
            ),
            'GetItem': grpc.unary_unary_rpc_method_handler(
                    servicer.GetItem,
                    request_deserializer=knowledge__graph__pb2.GetItemReq.FromString,
                    response_serializer=knowledge__graph__pb2.GetItemRes.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kg.KGService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class KGService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetGraph(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kg.KGService/GetGraph',
            knowledge__graph__pb2.GetGraphReq.SerializeToString,
            knowledge__graph__pb2.GetGraphRes.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetItem(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kg.KGService/GetItem',
            knowledge__graph__pb2.GetItemReq.SerializeToString,
            knowledge__graph__pb2.GetItemRes.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
