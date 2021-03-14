export interface VuexArgs<T> {
  state: T;
  commit: Function;
  dispatch: Function;
  getters: Record<string, any>;
}
