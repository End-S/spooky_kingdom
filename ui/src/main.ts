import { createApp } from "vue";
import { createPinia } from "pinia";
import Oruga from "@oruga-ui/oruga-next";

import App from "./App.vue";
import router from "./router";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { fas } from "@fortawesome/free-solid-svg-icons";
library.add(fas);

import "@oruga-ui/oruga-next/dist/oruga-full.css";
import "./assets/main.css";
import "./assets/oruga-tailwindcss.css";

const app = createApp(App).use(Oruga, {
  iconPack: "fas",
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
