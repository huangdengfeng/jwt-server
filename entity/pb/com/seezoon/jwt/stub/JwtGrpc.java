package com.seezoon.jwt.stub;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.64.0)",
    comments = "Source: jwt.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class JwtGrpc {

  private JwtGrpc() {}

  public static final java.lang.String SERVICE_NAME = "com.seezoon.jwt.Jwt";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.seezoon.jwt.stub.SignReq,
      com.seezoon.jwt.stub.SignResp> getSignMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Sign",
      requestType = com.seezoon.jwt.stub.SignReq.class,
      responseType = com.seezoon.jwt.stub.SignResp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.seezoon.jwt.stub.SignReq,
      com.seezoon.jwt.stub.SignResp> getSignMethod() {
    io.grpc.MethodDescriptor<com.seezoon.jwt.stub.SignReq, com.seezoon.jwt.stub.SignResp> getSignMethod;
    if ((getSignMethod = JwtGrpc.getSignMethod) == null) {
      synchronized (JwtGrpc.class) {
        if ((getSignMethod = JwtGrpc.getSignMethod) == null) {
          JwtGrpc.getSignMethod = getSignMethod =
              io.grpc.MethodDescriptor.<com.seezoon.jwt.stub.SignReq, com.seezoon.jwt.stub.SignResp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Sign"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.seezoon.jwt.stub.SignReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.seezoon.jwt.stub.SignResp.getDefaultInstance()))
              .setSchemaDescriptor(new JwtMethodDescriptorSupplier("Sign"))
              .build();
        }
      }
    }
    return getSignMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.seezoon.jwt.stub.VerifyReq,
      com.seezoon.jwt.stub.VerifyResp> getVerifyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Verify",
      requestType = com.seezoon.jwt.stub.VerifyReq.class,
      responseType = com.seezoon.jwt.stub.VerifyResp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.seezoon.jwt.stub.VerifyReq,
      com.seezoon.jwt.stub.VerifyResp> getVerifyMethod() {
    io.grpc.MethodDescriptor<com.seezoon.jwt.stub.VerifyReq, com.seezoon.jwt.stub.VerifyResp> getVerifyMethod;
    if ((getVerifyMethod = JwtGrpc.getVerifyMethod) == null) {
      synchronized (JwtGrpc.class) {
        if ((getVerifyMethod = JwtGrpc.getVerifyMethod) == null) {
          JwtGrpc.getVerifyMethod = getVerifyMethod =
              io.grpc.MethodDescriptor.<com.seezoon.jwt.stub.VerifyReq, com.seezoon.jwt.stub.VerifyResp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Verify"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.seezoon.jwt.stub.VerifyReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.seezoon.jwt.stub.VerifyResp.getDefaultInstance()))
              .setSchemaDescriptor(new JwtMethodDescriptorSupplier("Verify"))
              .build();
        }
      }
    }
    return getVerifyMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static JwtStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JwtStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JwtStub>() {
        @java.lang.Override
        public JwtStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JwtStub(channel, callOptions);
        }
      };
    return JwtStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static JwtBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JwtBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JwtBlockingStub>() {
        @java.lang.Override
        public JwtBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JwtBlockingStub(channel, callOptions);
        }
      };
    return JwtBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static JwtFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JwtFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JwtFutureStub>() {
        @java.lang.Override
        public JwtFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JwtFutureStub(channel, callOptions);
        }
      };
    return JwtFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * 签发
     * </pre>
     */
    default void sign(com.seezoon.jwt.stub.SignReq request,
        io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.SignResp> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSignMethod(), responseObserver);
    }

    /**
     * <pre>
     * 验证
     * </pre>
     */
    default void verify(com.seezoon.jwt.stub.VerifyReq request,
        io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.VerifyResp> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getVerifyMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Jwt.
   */
  public static abstract class JwtImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return JwtGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Jwt.
   */
  public static final class JwtStub
      extends io.grpc.stub.AbstractAsyncStub<JwtStub> {
    private JwtStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JwtStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JwtStub(channel, callOptions);
    }

    /**
     * <pre>
     * 签发
     * </pre>
     */
    public void sign(com.seezoon.jwt.stub.SignReq request,
        io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.SignResp> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSignMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 验证
     * </pre>
     */
    public void verify(com.seezoon.jwt.stub.VerifyReq request,
        io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.VerifyResp> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getVerifyMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Jwt.
   */
  public static final class JwtBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<JwtBlockingStub> {
    private JwtBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JwtBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JwtBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 签发
     * </pre>
     */
    public com.seezoon.jwt.stub.SignResp sign(com.seezoon.jwt.stub.SignReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 验证
     * </pre>
     */
    public com.seezoon.jwt.stub.VerifyResp verify(com.seezoon.jwt.stub.VerifyReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getVerifyMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Jwt.
   */
  public static final class JwtFutureStub
      extends io.grpc.stub.AbstractFutureStub<JwtFutureStub> {
    private JwtFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JwtFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JwtFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 签发
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.seezoon.jwt.stub.SignResp> sign(
        com.seezoon.jwt.stub.SignReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSignMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 验证
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.seezoon.jwt.stub.VerifyResp> verify(
        com.seezoon.jwt.stub.VerifyReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getVerifyMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SIGN = 0;
  private static final int METHODID_VERIFY = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SIGN:
          serviceImpl.sign((com.seezoon.jwt.stub.SignReq) request,
              (io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.SignResp>) responseObserver);
          break;
        case METHODID_VERIFY:
          serviceImpl.verify((com.seezoon.jwt.stub.VerifyReq) request,
              (io.grpc.stub.StreamObserver<com.seezoon.jwt.stub.VerifyResp>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getSignMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.seezoon.jwt.stub.SignReq,
              com.seezoon.jwt.stub.SignResp>(
                service, METHODID_SIGN)))
        .addMethod(
          getVerifyMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.seezoon.jwt.stub.VerifyReq,
              com.seezoon.jwt.stub.VerifyResp>(
                service, METHODID_VERIFY)))
        .build();
  }

  private static abstract class JwtBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    JwtBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.seezoon.jwt.stub.JwtOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Jwt");
    }
  }

  private static final class JwtFileDescriptorSupplier
      extends JwtBaseDescriptorSupplier {
    JwtFileDescriptorSupplier() {}
  }

  private static final class JwtMethodDescriptorSupplier
      extends JwtBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    JwtMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (JwtGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new JwtFileDescriptorSupplier())
              .addMethod(getSignMethod())
              .addMethod(getVerifyMethod())
              .build();
        }
      }
    }
    return result;
  }
}
