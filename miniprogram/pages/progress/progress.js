// pages/progress/progress.js
const app = getApp()

Page({

  data: {
    progress: [],
    subjectType: "anime",
  },

  onLoad: function (options) {
    app.globalData.chii.userCollections(this.data.subjectType, function (res) {
      console.log(res.data);
    })
  },

})