import { createApp } from "vue";
import { createPinia } from "pinia";
import {
  Config,
  Notification,
  OButton,
  ODatepicker,
  ODropdown,
  OIcon,
  OInput,
  OField,
  ORadio,
  OModal,
  OPagination,
  ODropdownItem,
  OSkeleton,
  OInputitems,
} from "@oruga-ui/oruga-next";
import App from "./App.vue";
import router from "./router";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faPenToSquare,
  faFilter,
  faCircleHalfStroke,
  faNewspaper,
  faMoon,
  faSun,
  faCalendar,
  faCancel,
  faChevronLeft,
  faChevronRight,
  faWarning,
  faCaretDown,
  faBars,
  faEnvelope,
  faEye,
  faTimes,
} from "@fortawesome/free-solid-svg-icons";
library.add(
  faPenToSquare,
  faFilter,
  faCircleHalfStroke,
  faNewspaper,
  faMoon,
  faSun,
  faCalendar,
  faCancel,
  faChevronLeft,
  faChevronRight,
  faWarning,
  faCaretDown,
  faBars,
  faEnvelope,
  faEye,
  faTimes
);

import "@oruga-ui/oruga-next/dist/oruga-full.css";
import "./assets/main.css";
import "./assets/oruga-tailwindcss.css";

const app = createApp(App);
app.component("OButton", OButton);
app.component("ODatepicker", ODatepicker);
app.component("ODropdown", ODropdown);
app.component("OIcon", OIcon);
app.component("OInput", OInput);
app.component("OField", OField);
app.component("ORadio", ORadio);
app.component("OModal", OModal);
app.component("OPagination", OPagination);
app.component("ODropdownItem", ODropdownItem);
app.component("OSkeleton", OSkeleton);
app.component("OInputitems", OInputitems);
app.use(Notification);
app.use(Config, {
  iconPack: "fa",
  iconComponent: "vue-fontawesome",
  button: {
    override: true,
    rootClass: "btn",
  },
  switch: {
    checkClass: "switch",
    labelClass: "switch-label",
  },
  datepicker: {
    mobileModal: true,
    mobileNative: false,
    tableCellClass: "table-cell",
    tableCellUnselectableClass: "table-cell-unselectable",
    tableHeadCellClass: "table-head-cell",
    monthCellClass: "month-cell",
    monthCellUnselectableClass: "month-cell-unselectable",
    monthCellSelectedClass: "month-cell-selected",
    monthCellWithinSelectedClass: "month-cell-within-selected",
    monthCellFirstHoveredClass: "month-cell-first-hovered",
    monthCellLastHoveredClass: "month-cell-last-hovered",
    monthCellWithinHoveredClass: "month-cell-within-hovered",
    nextBtnClass: "next-btn",
    prevBtnClass: "prev-btn",
  },
  dropdown: {
    menuClass: "dropdown-menu",
    itemClass: "dropdown-item",
    itemActiveClass: "dropdown-item-active",
    triggerClass: "dropdown-trigger",
  },
  input: {
    inputClass: "input",
  },
  radio: {
    checkClass: "radio-check",
    checkCheckedClass: "radio-checked",
  },
  modal: {
    contentClass: "modal-content",
  },
  field: {
    labelClass: "field-label",
  },
  inputitems: {
    itemClass: "input-item",
  },
  skeleton: {
    itemClass: "skeleton",
    animationClass: "skeleton-animation",
  },
});

app.component("vue-fontawesome", FontAwesomeIcon);

app.use(createPinia());
app.use(router);

app.mount("#app");
