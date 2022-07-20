<template>
  <div>
    <el-table :data="tableData" style="width: 96%" height="100%" id="table">
      <el-table-column fixed prop="threatInfoID" label="威胁ID" width="120">
      </el-table-column>
      <el-table-column prop="ctime" label="时间" :formatter="formateTime">
      </el-table-column>
      <el-table-column prop="ioc" label="IP地址"> </el-table-column>
      <el-table-column prop="domainCount" label="解析域名" width="100">
      </el-table-column>
      <el-table-column prop="tagCount" label="标签数量" width="100">
      </el-table-column>
      <el-table-column prop="itelCount" label="线索数量" width="100">
      </el-table-column>
      <el-table-column prop="tagCount" label="标签数量" width="100">
      </el-table-column>
      <el-table-column
        prop="judge"
        label="类型"
        width="100"
        :formatter="formateJudgeType"
      >
      </el-table-column>
      <el-table-column prop="source" label="数据来源" width="200">
      </el-table-column>
    </el-table>

    <el-pagination
      background
      layout="prev, pager, next"
      :total="totalCount"
      :page-size="pageSize"
      id="slider"
      @current-change="onPageChange"
    >
    </el-pagination>
  </div>
</template>

<style>
#table {
  margin: 16px 2%;
}
#slider {
  padding-bottom: 32px;
}
</style>

<script>
export default {
  data() {
    return {
      data: [
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
        {
          threatInfoID: 14399,
          ioc: "211.95.50.4",
          domainCount: "0",
          tagCount: "5",
          itelCount: "10",
          judge: 2,
          poc: false,
          ctime: "1657694018004",
          source: 0,
        },
      ],
      totalCount: 126,
      pageSize: 10,
      page: 0,
      tableData: [],
    };
  },
  mounted: function () {
    this.getTableData();
  },
  methods: {
    getTableData: function () {
      var that = this;
      this.axios
        .get("https://baidu.com")
        .then((res) => {
          console.log(res.data);
          
        })
        .catch((err) => {
          console.log(err);
        });
    },
    formateJudgeType: function (row, column, cellValue, index) {
      var t = cellValue;
      var result = "——";
      switch (t) {
        case 0:
          result = "安全";
          break;
        case 1:
          result = "未知";
          break;
        case 2:
          result = "恶意";
          break;
        default:
          result = t;
          break;
      }
      return result;
    },
    formateTime: function (row, column, cellValue, index) {
      let time = new Date(parseInt(cellValue)).toLocaleString();
      return time;
    },
    interfaceTrans: function (rawData) {},
    onPageChange: function () {
      console.log("pageChange");
    },
  },
};
</script>