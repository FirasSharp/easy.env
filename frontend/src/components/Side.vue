<script setup>
import Button from "primevue/button";
import "primeicons/primeicons.css";
import Toolbar from "primevue/toolbar";
import Tree from "primevue/tree";
import { ref } from "vue";
import { FilePrompt, ListDatabases } from "../../wailsjs/go/main/App";
const nodes = ref();
const selectedKey = ref(null);
const emits = defineEmits(["envsChanged"]);

nodes.value = [
  {
    key: "path/to/file",
    label: "Test",
    icon: "pi pi-fw pi-database",
    children: [
      {
        key: "project",
        label: "Projects",
        icon: "pi pi-fw pi-file",
        children: [
          {
            key: "51c6da3b-af53-4e46-bca4-b36ba2cb96b8",
            label: "exampleProject12323",
            icon: "pi pi-fw pi-table",
            path: `C:\\Users\\Firas\\Desktop\\git\\easy.env-lib\\_examples\\project`,
            envs: [
              { key: "aaa", value: "wwwe" },
              { key: "aaa213", value: "wwwe3424" },
              { key: "aaatt", value: "www23432e" },
            ],
          },
        ],
      },
    ],
  },
];
const onNodeSelect = (e) => {
  emits("envsChanged", [
    { key: "aaa", value: "wwwe" },
    { key: "aaa213", value: "wwwe3424" },
    { key: "aaatt", value: "www23432e" },
  ]);
  console.log("123");
};
const onOpenDB = async () => {
  var ok = await FilePrompt();
  if (ok) {
    var db = ListDatabases();
    console.log(db)
  } else {
    debugger;
  }
};
</script>
<template>
  <Toolbar>
    <template #start>
      <div>Databases</div>
    </template>

    <template #end>
      <Button
        aria-label="Open database"
        aria-placeholder="Open database"
        placeholder="Open database"
        icon="pi pi-folder-open"
        severity="secondary"
        text
        outlined
        size="small"
        @click="onOpenDB"
      />
    </template>
  </Toolbar>
  <Tree
    id="tree"
    class="noBackground"
    v-model:selectionKeys="selectedKey"
    :value="nodes"
    @nodeSelect="onNodeSelect"
    selectionMode="single"
  >
  </Tree>
</template>

<style scoped>
.noBackground {
  background-color: none;
}
</style>
