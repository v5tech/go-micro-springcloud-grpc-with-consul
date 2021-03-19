package com.anchnet.grpc.grpc.server;


import com.anchnet.grpc.proto.HelloGrpc;
import com.anchnet.grpc.proto.HelloReply;
import com.anchnet.grpc.proto.HelloRequest;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

@GrpcService
public class GrpcServerService extends HelloGrpc.HelloImplBase {
    @Override
    public void sayHello(HelloRequest request, StreamObserver<HelloReply> responseObserver) {
        HelloReply reply = HelloReply.newBuilder().setMessage("Hello ==> " + request.getName()).build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
