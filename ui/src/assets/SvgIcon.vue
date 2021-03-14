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
  EyeIcon,
  EyeOffIcon,
  FilterIcon,
  MenuIcon,
  MinusSquareIcon,
  PlusSquareIcon,
  InfoIcon,
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
    EyeIcon,
    EyeOffIcon,
    FilterIcon,
    InfoIcon,
  },
  data() {
    return {
      eye: EyeIcon,
      'eye-off': EyeOffIcon,
      information: InfoIcon,
    };
  },
  computed: {
    iconName(this: any) {
      const icon = this.$props.icon[ 1 ];
      // not all buefy components support custom icons
      // translate hard-coded buefy icons to our custom icons
      if (this.$data[ icon ]) return this.$data[ icon ];

      return camelCase(icon);
    },
    iconProps() {
      return { size: this.$props.size || '1x' };
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
