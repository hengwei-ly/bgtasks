ITSM： IT服务管理软件


数据表和对应的访问路径：
基础表：
http://localhost:9000/hengwei/aaa/Companies/Index  		            公司信息
http://localhost:9000/hengwei/aaa/Departments/Index		            部门信息
http://localhost:9000/hengwei/aaa/Employees/Index		            员工信息
http://localhost:9000/hengwei/aaa/Users/Index			            登录信息
http://localhost:9000/hengwei/aaa/Slas/Index			            SLA 服务登记协议
http://localhost:9000/hengwei/aaa/Cabs/Index			            CAB 变更管理委员会
http://localhost:9000/hengwei/aaa/CabEmployees/Index                变更管理委员会对应的组员
http://localhost:9000/hengwei/aaa/RequestTypes/Index		        请求类型
http://localhost:9000/hengwei/aaa/Systemproperties/Index	        全局属性，部分比较简单的属性都放在这个表里面{name:value}
http://localhost:9000/hengwei/aaa/SupportGroups/Index		        支持组
http://localhost:9000/hengwei/aaa/SupportGroupEmployees/Index	    支持组-组员


解决方案：
http://localhost:9000/hengwei/aaa/Solutions/Index 		            解决方案
http://localhost:9000/hengwei/aaa/SolutionTopics/Index		        方案主题

审批：
http://localhost:9000/hengwei/aaa/Approvals/Index  	                审批

请求：
http://localhost:9000/hengwei/aaa/Requests/Index		            请求信息，多值显示的部分还没有完成
http://localhost:9000/hengwei/aaa/Requesters/Index                  请求人信息，与员工不同，表示的是公司外发出请求的人员
http://localhost:9000/hengwei/aaa/Categories/Index                  类别，树状表示
http://localhost:9000/hengwei/aaa/RequestAssets/Index		        请求-资产

变更：
http://localhost:9000/hengwei/aaa/Changes/Index			            变更信息
http://localhost:9000/hengwei/aaa/ChangestageChangeactions/Index    变更阶段-变更动作
http://localhost:9000/hengwei/aaa/ChangestageStatuses/Index	        变更阶段-变更状态
http://localhost:9000/hengwei/aaa/Changestages/Index		        变更状态
http://localhost:9000/hengwei/aaa/Changeactions/Index		        变更动作
http://localhost:9000/hengwei/aaa/Changeroles/Index                 变更角色（应该包含CAB）
http://localhost:9000/hengwei/aaa/ChangeactionChangeroles/Index	    变更动作-变更角色
http://localhost:9000/hengwei/aaa/ChangestageStatusEmployees/Index  变更阶段的状态和对应要通知到的人员

  变更工作流：
http://localhost:9000/hengwei/aaa/Changeworkflowurgencies/Index     变更工作流的紧急度
http://localhost:9000/hengwei/aaa/Changeworkflows/Index		        变更工作流
http://localhost:9000/hengwei/aaa/ChangeWorkflowChagestages/Index   变更工作流-变更状态

  变更关联：
http://localhost:9000/hengwei/aaa/ChangeAssets/Index                变更-资产

问题：
http://localhost:9000/hengwei/aaa/Problems/Index		            问题
http://localhost:9000/hengwei/aaa/ProblemAssets/Index               问题-资产

项目：
http://localhost:9000/hengwei/aaa/ProjectroleProjectactions/Index   项目-项目动作
http://localhost:9000/hengwei/aaa/Projectroles/Index		        项目角色
http://localhost:9000/hengwei/aaa/Projectactions/Index		        项目动作
http://localhost:9000/hengwei/aaa/Projects/Index		            项目信息
http://localhost:9000/hengwei/aaa/ProjectEmployees/Index	        项目-项目参与人

里程碑
http://localhost:9000/hengwei/aaa/Milestones/Index		            里程碑信息，这里是只做一级

其他：包括历史、任务、注释、工作日志等
http://localhost:9000/hengwei/aaa/Histories/Index		            历史信息
http://localhost:9000/hengwei/aaa/Tasks/Index			            任务信息
http://localhost:9000/hengwei/aaa/Annotations/Index		            注释信息
http://localhost:9000/hengwei/aaa/Worklogs/Index		            工作日志信息

