package org.bal.frontend.config;

import brave.grpc.GrpcTracing;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.health.v1.HealthGrpc;
import org.bal.quote.proto.internal.QuoteManagementGrpc;
import org.bal.vote.proto.internal.VoteManagementGrpc;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class VoteServiceConfiguration {

    @Value("${vote-server.host}")
    private String voteServiceHost;

    @Value("${vote-server.port}")
    private int voteServicePort;

    @Autowired
    private GrpcTracing grpcTracing;

    @Bean
    public ManagedChannelBuilder voteServiceManagedChannelBuilder() {
        return ManagedChannelBuilder.forAddress(voteServiceHost, voteServicePort).intercept(grpcTracing.newClientInterceptor())
                .usePlaintext(true);
    }

    @Bean
    public ManagedChannel voteServiceManagedChannel() {
        return voteServiceManagedChannelBuilder().build();
    }

    @Bean("voteManagementBlockingStub")
    public VoteManagementGrpc.VoteManagementBlockingStub voteManagementBlockingStub() {
        return VoteManagementGrpc.newBlockingStub(voteServiceManagedChannel());
    }

    @Bean("voteHealthBlockingStub")
    public HealthGrpc.HealthBlockingStub voteHealthBlockingStub() {
        return HealthGrpc.newBlockingStub(voteServiceManagedChannel());
    }

}
