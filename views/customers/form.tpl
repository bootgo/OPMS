<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 客户管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/customer/list">客户管理</a> </li>
        <li class="active"> 客户 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="customer-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="name" value="{{.customer.Name}}" class="form-control" placeholder="请输入客户名称">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">联系人</label>
                  <div class="col-sm-10">
                    <input type="text" name="contact" value="{{.customer.Contact}}" class="form-control" placeholder="请输入联系人">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">联系地址</label>
                  <div class="col-sm-10">
                    <input type="text" name="address" value="{{.customer.Address}}" class="form-control" placeholder="请输入客户地址">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">办公电话</label>
                  <div class="col-sm-10">
                    <input type="text" name="tel" value="{{.customer.Tel}}" class="form-control" placeholder="请输入办公电话">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">手机</label>
                  <div class="col-sm-10">
                    <input type="text" name="phone" value="{{.customer.Phone}}" class="form-control" placeholder="请输入手机号">
                  </div>
                </div>
				
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">Email</label>
                  <div class="col-sm-10">
                    <input type="text" name="email" value="{{.customer.Email}}" class="form-control" placeholder="请输入Email">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">QQ</label>
                  <div class="col-sm-10">
                    <input type="text" name="qq" value="{{.customer.Qq}}" class="form-control" placeholder="请输入QQ号">
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">微信</label>
                  <div class="col-sm-10">
                    <input type="text" name="wechat" value="{{.customer.Wechat}}" class="form-control" placeholder="请输入微信号">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">备注</label>
                  <div class="col-sm-10">
                    <textarea name="note" placeholder="备注说明" style="height:300px;" class="form-control">{{.customer.Note}}</textarea>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户附件</label>
                  <div class="col-sm-10">
                    <input type="file" name="attachment">
                    {{if ne .customer.Attachment ""}}<br/>
                    <a href="{{.customer.Attachment}}" target="_blank">预览下载</a> {{end}} </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">状态</label>
                  <div class="col-sm-10">
                    <label class="radio-inline">
                    <input type="radio" name="status" value="1" {{if eq 1 .customer.Status}}checked{{end}}>
                    正常 </label>
                    <label class="radio-inline">
                    <input type="radio" name="status" value="2" {{if eq 2 .customer.Status}}checked{{end}}>
                    禁用 </label>
                    <label class="radio-inline">
                    <input type="radio" name="status" value="3" {{if eq 3 .customer.Status}}checked{{end}}>
                    待评估 </label>
                    <label class="radio-inline">
                    <input type="radio" name="status" value="4" {{if eq 4 .customer.Status}}checked{{end}}>
                    洽谈中 </label>
                    <label class="radio-inline">
                    <input type="radio" name="status" value="5" {{if eq 5 .customer.Status}}checked{{end}}>
                    不可用 </label>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" id="customerid" value="{{.customer.Id}}">
                    <button type="submit" class="btn btn-primary">提交保存</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script src="/static/js/datepicker-zh-CN.js"></script>
<script src="/static/keditor/kindeditor-min.js"></script>
<script>
$(function(){
	var editor = KindEditor.create('textarea[name="note"]', {
	    uploadJson: "/kindeditor/upload",
	    allowFileManager: true,
	    filterMode : false,
	    afterBlur: function(){this.sync();}
	});

 	$('#default-date-picker').datepicker('option', $.datepicker.regional['zh-CN']); 	
	 $('#default-date-picker').datepicker({
        dateFormat: 'yy-mm-dd',
		changeMonth: true,
		changeYear: true,
		yearRange:'-60:+0'
    });
})
</script>
</body>
</html>
