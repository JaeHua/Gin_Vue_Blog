<template>
  <router-view />
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  })
}
</script>
<style>
/* 滚动条样式 */
::-webkit-scrollbar {
  width: 12px;
}

::-webkit-scrollbar-track {
  background: #ffe4e1; /* 浅粉色背景 */
}

::-webkit-scrollbar-thumb {
  background-color: #ffb6c1; /* 粉红色滚动条 */
  border-radius: 10px;
  border: 3px solid #ffe4e1;
}

::-webkit-scrollbar-thumb:hover {
  background: #ff69b4; /* 浅紫红色 */
}

* {
  scrollbar-width: thin;
  scrollbar-color: #ffb6c1 #ffe4e1;
}
</style>
