import { isErrorResponse } from "@/common/models/api.model";
import { debounce, type DebouncedFunc } from "lodash-es";
import { defineStore } from "pinia";
import { useProgrammatic } from "@oruga-ui/oruga-next";

interface GlobalState {
  err: string;
  activeRequests: number;
  showLoadingActivity: boolean;
}

const debounceLoading: DebouncedFunc<() => void> = debounce(() => {
  const globalStore = useGlobalStore();
  globalStore.showLoadingActivity = globalStore.activeRequests > 0;
}, 500);

export const useGlobalStore = defineStore("globalStore", {
  state: (): GlobalState => ({
    err: "",
    activeRequests: 0,
    showLoadingActivity: false,
  }),
  actions: {
    incrementActiveRequests() {
      this.activeRequests += 1;
      // only activate loading state if active request persists
      debounceLoading();
    },
    decrementActiveRequests() {
      this.activeRequests -= 1;
      if (this.activeRequests === 0) {
        this.showLoadingActivity = false;
      }
    },
    error(err: unknown) {
      const { oruga } = useProgrammatic();
      console.error(err);
      let message = "Something went wrong";
      if (isErrorResponse(err) && err.error) {
        message = `${err.error} (${err.status})`;
      }
      this.err = message;
      oruga.notification.open({
        type: "warning",
        variant: "warning",
        duration: 5000,
        message,
        position: "bottom",
        queue: true,
        closable: true,
      });

      throw err;
    },
  },
});
