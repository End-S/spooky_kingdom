<template>
  <!--  dynamic component used to display the icon component -->
  <component :is="iconName" v-bind="iconProps"></component>
</template>

<script lang="ts">

import {
  Component,
  Prop,
  Vue,
} from 'vue-property-decorator';
import { camelCase } from 'lodash-es';
import {
  CalendarIcon,
  ChevronDownIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  MenuIcon,
  MinusSquareIcon,
  PlusSquareIcon,
} from 'vue-feather-icons';

@Component({
  components: {
    ChevronRightIcon,
    ChevronLeftIcon,
    CalendarIcon,
    PlusSquareIcon,
    MinusSquareIcon,
    MenuIcon,
    ChevronDownIcon,
  },
  computed: {
    iconName() {
      return camelCase(this.$props.icon[ 1 ]);
    },
    iconProps() {
      return { size: this.$props.size || '2x' };
    },
  },
})
export default class SvgIcon extends Vue {
  // buefy passes: [ icon pack, icon name ] to this prop
  @Prop() private icon!: [ string, string ];
  // buefy passes customSize to this prop
  @Prop() private size!: string;
}

</script>

<style lang="scss">

.icon svg {
  stroke-width: 2 !important;
  fill: none !important;
}
</style>
