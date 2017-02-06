<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <form class="searchform" action="/customer/list" method="get">
        <select name="status" class="form-control">
          <option value="">状态</option>
          <option value="1" {{if eq "1" .condArr.status}}selected{{end}}>正常</option>
          <option value="2" {{if eq "2" .condArr.status}}selected{{end}}>禁用</option>
		<option value="3" {{if eq "3" .condArr.status}}selected{{end}}>待评估</option>
		<option value="4" {{if eq "4" .condArr.status}}selected{{end}}>洽谈中</option>
		<option value="5" {{if eq "5" .condArr.status}}selected{{end}}>不可用</option>
        </select>
        <input type="text" class="form-control" name="keywords" placeholder="请输入客户名称" value="{{.condArr.keywords}}"/>
        <button type="submit" class="btn btn-primary">搜索</button>
      </form>
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 客户管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/customer/list">客户管理</a> </li>
        <li class="active"> 客户 </li>
      </ul>
      <div class="pull-right"><a href="/customer/add" class="btn btn-success">添加新客户</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 客户管理 / 总数：{{.countCustomer}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="customer-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>客户名称</th>
                        <th>联系人</th>
						<th>地址</th>
						<th>座机</th>
						<th>手机</th>
						<th>Email</th>
						<th>QQ</th>
						<th>微信</th>
						<th>附件</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .customers}}
                    <tr>
                      <td>{{$v.Name}}</td>
					  <td>{{$v.Contact}}</td>
                      <td>{{$v.Address}}</td>
					<td>{{$v.Tel}}</td>
					<td>{{$v.Phone}}</td>
					<td>{{$v.Email}}</td>
					<td>{{$v.Qq}}</td>
					<td>{{$v.Wechat}}</td>
					<td><a href="{{$v.Attachment}}" target="_blank">查看预览</a></td>
                      <td>{{getCustomerStatus $v.Status}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/customer/edit/{{$v.Id}}">编辑</a></li>                           
                            {{if eq 1 $v.Status}}
							 <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-customers-single" data-id="{{$v.Id}}" data-status="2">禁用</a></li>
                            {{else if eq 2 $v.Status}}
							 <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-customers-single" data-id="{{$v.Id}}" data-status="3">待评估</a></li>
							<li><a href="javascript:;" class="js-customers-single" data-id="{{$v.Id}}" data-status="4">洽谈中</a></li>                           
							<li><a href="javascript:;" class="js-customers-single" data-id="{{$v.Id}}" data-status="5">不可用</a></li>                            
                            {{end}}
							<li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-customers-delete" data-id="{{$v.Id}}">删除</a></li>
                          </ul>
                        </div></td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
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
</body>
</html>
