import { HTTP } from "@/common/http";
import type { Publisher } from "@/common/models/publisher.model";
import type { ApiResponse } from "@/common/models/api.model";

export type GetPublishersRes = ApiResponse<{ publishers: Publisher[] }>;
export const get = (): Promise<GetPublishersRes> => HTTP.get("publishers");
