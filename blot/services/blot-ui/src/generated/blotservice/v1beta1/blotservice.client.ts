// @generated by protobuf-ts 2.9.4 with parameter use_proto_field_name
// @generated from protobuf file "blotservice/v1beta1/blotservice.proto" (package "blotservice.v1beta1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { BlotService } from "./blotservice";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetGameForPlayerResponse } from "./blotservice";
import type { GetGameForPlayerRequest } from "./blotservice";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service blotservice.v1beta1.BlotService
 */
export interface IBlotServiceClient {
    /**
     * @generated from protobuf rpc: GetGameForPlayer(blotservice.v1beta1.GetGameForPlayerRequest) returns (blotservice.v1beta1.GetGameForPlayerResponse);
     */
    getGameForPlayer(input: GetGameForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameForPlayerRequest, GetGameForPlayerResponse>;
}
/**
 * @generated from protobuf service blotservice.v1beta1.BlotService
 */
export class BlotServiceClient implements IBlotServiceClient, ServiceInfo {
    typeName = BlotService.typeName;
    methods = BlotService.methods;
    options = BlotService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetGameForPlayer(blotservice.v1beta1.GetGameForPlayerRequest) returns (blotservice.v1beta1.GetGameForPlayerResponse);
     */
    getGameForPlayer(input: GetGameForPlayerRequest, options?: RpcOptions): UnaryCall<GetGameForPlayerRequest, GetGameForPlayerResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetGameForPlayerRequest, GetGameForPlayerResponse>("unary", this._transport, method, opt, input);
    }
}