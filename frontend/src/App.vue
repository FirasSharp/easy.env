<script setup>
import Splitter from "primevue/splitter";
import SplitterPanel from "primevue/splitterpanel";
import Side from "./components/Side.vue";
import Main from "./components/Main.vue";
import { ref } from "vue";

const envs = ref();
envs.value = [];

const onEnvsChanged = (e) => {
  envs.value = e;
};
</script>

<template>
  <Splitter class="full-height" :gutterSize="1">
    <SplitterPanel class="side" :size="25"
      ><Side @envsChanged="onEnvsChanged"
    /></SplitterPanel>
    <SplitterPanel class="main" :size="70"><Main :envs="envs" /></SplitterPanel>
  </Splitter>
</template>

<style>
.full-height {
  height: 100%;
}
.side {
  background-color: #161618;
}
.main {
  background-color: #212124;
}
</style>
