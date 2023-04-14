import { get, type GetPublishersRes } from "@/api/publishers";
import { isErrorResponse } from "@/common/models/api.model";
import type { Publisher } from "@/common/models/publisher.model";
import { defineStore } from "pinia";
import { useGlobalStore } from "./global.store";

interface PublisherState {
  publishers: Publisher[];
}

export const usePublisherStore = defineStore("publisherStore", {
  state: (): PublisherState => ({
    publishers: [],
  }),
  getters: {
    getPublisherLabel: (state: PublisherState) => (id: string) => {
      const match = state.publishers.find((p) => p.id === id)?.label;
      return match || "?";
    },
  },
  actions: {
    setPublishers(publishers: Publisher[]) {
      this.publishers = publishers;
    },
    async getPublishers(): Promise<GetPublishersRes> {
      const res = await get();
      if (isErrorResponse(res)) {
        const globalStore = useGlobalStore();
        globalStore.error(res);
      } else {
        this.setPublishers(res?.publishers || []);
      }
      return res;
    },
  },
});
