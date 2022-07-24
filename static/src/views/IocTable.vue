<template>
  <div>
    <el-table
      :data="tableData"
      style="width: 96%"
      min-height="256px"
      id="table"
      :cell-class-name="tableCellClassName"
      @row-click="onRowClick"
    >
      <el-table-column fixed prop="threat_id_info" label="威胁ID" width="120">
        <template slot-scope="scope">
          <span style="cursor: pointer">
            {{ scope.row.threat_id_info }}
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="ctime" label="时间" :formatter="formateTime">
      </el-table-column>
      <el-table-column prop="ip" label="IP地址"> </el-table-column>
      <el-table-column prop="domain_count" label="解析域名" width="100" :formatter="dataWasher">
      </el-table-column>
      <el-table-column prop="tag_count" label="标签数量" width="100" :formatter="dataWasher">
      </el-table-column>
      <el-table-column prop="itel_count" label="线索数量" width="100" :formatter="dataWasher">
      </el-table-column>
      <el-table-column
        prop="judge"
        label="类型"
        width="100"
        :formatter="formateJudgeType"
      >
      </el-table-column>
      <el-table-column prop="source" label="数据来源" width="200" :formatter="formateSource">
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
          id: 93,
          ip: "14.29.118.14",
          threat_id_info: "14556",
          tag_count: 3,
          itel_count: 6,
          judge: 2,
          source: 1,
          ctime: 1657790439673,
        },
        {
          id: 109,
          ip: "157.245.78.71",
          threat_id_info: "14556",
          tag_count: 4,
          itel_count: 7,
          judge: 2,
          source: 1,
          ctime: 1657790439673,
        },
        {
          id: 14,
          ip: "101.86.144.28",
          threat_id_info: "14661",
          domain_count: 1,
          tag_count: 1,
          itel_count: 5,
          judge: 1,
          source: 1,
          ctime: 1657952606068,
        },
        {
          id: 145,
          ip: "129.146.65.82",
          threat_id_info: "14420",
          tag_count: 7,
          itel_count: 9,
          judge: 2,
          source: 1,
          ctime: 1657639280716,
        },
        {
          id: 87,
          ip: "120.86.236.18",
          threat_id_info: "14576",
          tag_count: 2,
          itel_count: 3,
          judge: 2,
          source: 1,
          ctime: 1657848135824,
        },
        {
          id: 133,
          ip: "112.96.80.229",
          threat_id_info: "14429",
          tag_count: 1,
          itel_count: 3,
          judge: 1,
          source: 1,
          ctime: 1657676372732,
        },
        {
          id: 90,
          ip: "106.75.129.215",
          threat_id_info: "14568",
          domain_count: 5,
          tag_count: 5,
          itel_count: 8,
          judge: 2,
          source: 1,
          ctime: 1657840523147,
        },
        {
          id: 40,
          ip: "14.18.81.80",
          threat_id_info: "14556",
          domain_count: 1,
          tag_count: 3,
          itel_count: 6,
          judge: 2,
          source: 1,
          ctime: 1657790439673,
        },
        {
          id: 95,
          ip: "64.227.188.188",
          threat_id_info: "14556",
          tag_count: 3,
          itel_count: 8,
          judge: 2,
          source: 1,
          ctime: 1657790439673,
        },
        {
          id: 20,
          ip: "174.138.63.58",
          threat_id_info: "14626",
          domain_count: 83,
          tag_count: 3,
          itel_count: 6,
          judge: 2,
          source: 1,
          ctime: 1657882515888,
        },
      ],
      totalCount: 0,
      pageSize: 10,
      page: 1,
      tableData: [],
    };
  },
  mounted: function () {
    this.getTableData();
  },
  methods: {
    getTableData: function () {
      var that = this;
      this.axios({
        url: "threats",
        method: "GET",
        params: {
          current: this.page,
          page_size: this.pageSize,
        },
      })
        .then((res) => {
          console.log(res.data);
          if (res.data.code == 0) {
            that.data = res.data;
            that.totalCount = res.data.data.pagination.total;
            that.tableData = res.data.data.list;
            console.log(that.totalCount);
            //that.$refs.table.doLatout();
          } else {
            console.log("网络错误");
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    formateJudgeType: function (row, column, cellValue, index) {
      var t = cellValue;
      let result = "-"
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
          result = "-";
          break;
      }
      return result;
    },
    formateTime: function (row, column, cellValue, index) {
      let time = new Date(parseInt(cellValue)).toLocaleString();
      return time;
    },
    formateSource:function (row, column, cellValue, index){
      var t = cellValue;
      console.log(t)
      let result = "-"
      switch (t) {
        case -1:
          result = "爬取";
          break;
        case 1:
          result = "IOC";
          break;
        default:
          result = "-";
          break;
      }
      return result;
    },
    cellClick: function (row, column, cell, event) {
      /*
	    console.log(row)
	    console.log(column)
	    console.log(cell)
      console.log(event)
      */
    },
    tableCellClassName({ row, column, rowIndex, columnIndex }) {
      //注意这⾥是解构
      //利⽤单元格的 className 的回调⽅法，给⾏列索引赋值
      row.index = rowIndex;
      column.index = columnIndex;
    },
    onRowClick(row, column) {
      if (column.index == 0) {
        let index = row.index;
        let id = this.tableData[index].threat_id_info;
        let url = "https://x.threatbook.com/v5/article?threatInfoID=" + id;
        window.open(url, "_blank");
      }
    },
    interfaceTrans: function (rawData) {},
    onPageChange: function (p) {
      console.log(p);
      this.page = p;
      this.getTableData();
    },
    dataWasher: function (row, column, cellValue, index) {
      if (cellValue == -1 || cellValue == undefined) {
        return "-";
      } else {
        return cellValue;
      }
    },
  },
};
</script>