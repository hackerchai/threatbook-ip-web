<template>
  <div id="app">
    <div id="nav">
      <img src="./assets/logo.png" class="logo">
      <h2>威胁情报管理系统</h2>
    </div>
    <router-view />
  </div>
</template>

<style>
* {
  padding: 0;
  margin: 0;
}

:root {
  --color: #fff;
  --color-heighter: #05f2f2;
  --color-bg-main: #040c1f;
  --color-bg-api: #040c1f;
  --color-bg-hover: #0d2738;
  --color-primary: #158876;
  --color-secondary: #0e2233;
  --bg-secondary: #fff;
  --border-color: #ddd;
  --color-query: #05f2f2;
  --color-hover-query: #009688;
  --color-search-history-border: #05f2f2;
  --color-websit: #a7bcd1;
  --color-bg-websit: #041527;
  --color-bg-websit-heighter: #040c1f;
  --color-website-text: #05f2f2;
  --color-autocomplate-border: #05f2f2;
  --color-gray: #a7bcd1;
  --color-autocomplate-bg: #040c1f;
  --color-scrollbar-thumb: #119098;
  --collapse-bg: #061e2e;
  --collapse-bg-hover: #055360;
  --collapse-color: #05f2f2;
  --collaps-wrap-bg: #061e2e;
  --collaps-wrap-color: #fff;
  --dropdown-menu-bg: hsla(0, 0%, 100%, 0.06);
  --dropdown-menu-color: #fff;
  --poper-arrow-color: #061e2e;
  --user-bg: rgba(5, 242, 242, 0.02);
  --user-info-name-color: #fff;
  --user-nav-color: #a7bcd1;
  --person-center-bg: rgba(5, 242, 242, 0.02);
  --input-border-color: #05f2f2;
  --table-cell-hover: #0c2638;
  --table-cell-border: #0c2638;
  --table-cell-color: #fff;
  --dropdown-bg: #041024;
  --dropdown-hover-bg: #0d2738;
  --dropdown-border: #05f2f2;
  --dropdown-color: #fff;
  --dropdown-checked-color: #05f2f2;
  --button-confirm-color: #041527;
  --button-cancel-color: #fff;
  --button-cancel-bg: #061e2e;
  --person-line: #061e2e;
  --person-color: #606266;
  --person-label-color: #a7bcd1;
  --serarch-history-color: #a7bcd1;
}

#app {
  text-align: center;
  background-color: var(--color-bg-main);
  min-height: 100vh;
}

#nav {
  background-color: var(--brand-color);
  height: 64px;
  color: var(--color-heighter);
  display: flex;
  align-items: center;
}
.logo{
  height: 48px;
  width: 48px;
  padding: 0;
  margin:  0 2%;
}
</style>
