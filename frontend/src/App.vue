<script setup>
import Main from "./components/Main.vue";
</script>

<template>
  <div class="flex min-h-screen">
    <div id="SideNav" class="w-64 bg-gray-50 border-r border-gray-200">
      <div class="py-4 px-6"></div>

      <div class="mb-10">
        <h3 class="mx-6 mb-2 text-xs text-gray-400 uppercase tracking-widest">
          Main
        </h3>

        <!--Main component!-->
      </div>

      <div class="mb-10">
        <h3 class="mx-6 mb-2 text-xs text-gray-400 uppercase tracking-widest">
          Projects
        </h3>

        <!--listed projects!-->
      </div>
      <div class="mb-10">
        <h3 class="mx-6 mb-2 text-xs text-gray-400 uppercase tracking-widest">
          Templates
        </h3>

        <!--listed templates!-->
      </div>
    </div>
    <div class="flex-1">
      <main>
        <div class="inline-flex rounded-md shadow-sm" role="group">
          <div class="relative overflow-x-auto">
            <table
              class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
            >
              <thead
                class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
              >
                <tr>
                  <th scope="col" class="px-6 py-3">Key</th>
                  <th scope="col" class="px-6 py-3">Value</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                >
                  <td class="px-6 py-4">Silver</td>
                  <td class="px-6 py-4">Silver</td>
                </tr>
                <tr
                  class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                >
                  <td class="px-6 py-4">White</td>
                  <td class="px-6 py-4">Laptop PC</td>
                </tr>
                <tr class="bg-white dark:bg-gray-800">
                  <td class="px-6 py-4">Black</td>
                  <td class="px-6 py-4">Accessories</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style>
#SideNav {
  background-color: #202020;
}
.draggable {
  --wails-draggable: drag;
}
</style>
