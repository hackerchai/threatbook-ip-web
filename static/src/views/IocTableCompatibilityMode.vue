<template>
  <div>
    <el-table :data="tableData" style="width: 100%" height="250" id="table">
      <el-table-column fixed prop="ip" label="IP地址" width="150">
      </el-table-column>
      <el-table-column prop="time" label="时间" width="600"> </el-table-column>
    </el-table>

    <el-pagination background layout="prev, pager, next" :total="1000">
    </el-pagination>
  </div>
</template>

<style>
#table {
  width: 100%;
}
</style>

<script>
export default {
  data() {
    return {
      tableData: [
        {
          ip:"1.2.3.4",
          time:"2022-07-12 12:32:19"
        },
      ],
    };
  },
};
</script>