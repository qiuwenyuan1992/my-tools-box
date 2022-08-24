<template>
  <el-container style="height:100%">
    <el-aside style="width:auto;">
      <com-nav-menu style="height:100vh;"></com-nav-menu>
    </el-aside>
    <el-container >
      <el-header>Header</el-header>
      <el-main>
        <!-- 引入路由view -->
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>
<script lang="ts">

export default {
  name: "Home"
}
</script>