"""
You have to set up a server to run these tests:
```
python server/app.py
```
"""
import pytest
from grpc_requests import StubClient

from proto.knowledge_graph_pb2 import DESCRIPTOR, GetGraphRes, GetItemRes


class TestKnowledgeGraphService:

    @staticmethod
    def set_client() -> StubClient:
        service_descriptor = DESCRIPTOR.services_by_name["KGService"]
        return StubClient.get_by_endpoint("localhost:5555", service_descriptors=[service_descriptor])

    @pytest.mark.int
    def test_services(self):
        client = self.set_client()
        assert client.service_names == ["kg.KGService"]

    @pytest.mark.int
    def test_GetGraph_grpc(self):
        client = self.set_client()
        res = client.request("kg.KGService", "GetGraph",
                             {"description": "test"},
                             raw_output=True)
        assert isinstance(res.__class__, type(GetGraphRes))

    @pytest.mark.int
    def test_GetItem_grpc(self):
        client = self.set_client()
        res = client.request("kg.KGService", "GetItem",
                             {"id": "4e10c22261836cd0c415d2731465d43e"},
                             raw_output=True)
        assert isinstance(res.__class__, type(GetItemRes))
