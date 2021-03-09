## git 规范

1. 最稳定的代码放在 master 分支上（相当于 SVN 的 trunk 分支），我们不要直接在 master 分支上提交代码，只能在该分支上进行代码合并操作，例如将其它分支的代码合并到 master 分支上。
   
2. 我们日常开发中的 代码需要从 master 分支拉一条 develop 分支出来，该分支所有人都能访问，但一般情况下，我们也不会直接在该分支上提交代码，代码同样是从其它分支合并到 develop 分支上去。
   
### 分支管理规范

    master分支
    
    - 主分支，永远处于稳定状态，对应当前线上版本
    - 以 tag 标记一个版本，因此在 master 分支上看到的每一个 tag 都应该对应一个线上版本
    - 不允许在该分支直接提交代码
    
    develop 分支
    
    - 开发分支，包含了项目最新的功能和代码，所有开发都依赖 develop 分支进行
    - 小的改动可以直接在 develop 分支进行，改动较多时切出新的 feature 分支进行
    
    release 分支
    
    - 发布分支，新功能合并到 develop 分支，准备发布新版本时使用的分支
    - 当 develop 分支完成功能合并和部分 bug fix，准备发布新版本时，切出一个 release 分支，来做发布前的准备，分支名约定为release/xxx
    - 发布之前发现的 bug 就直接在这个分支上修复，确定准备发版本就合并到 master 分支，完成发布，同时合并到 develop 分支
    
    hotfix 分支
    
    - 紧急修复线上 bug 分支
    - 当线上版本出现 bug 时，从 master 分支切出一个 hotfix/xxx 分支，完成 bug 修复，然后将 hotfix/xxx 合并到 master 和 develop 分支(如果此时存在 release 分支，则应该合并到 release 分支)，合并完成后删除该 hotfix/xxx 分支
    
    master 分支: 线上稳定版本分支
    develop 分支: 开发分支，衍生出 feature 分支和 release 分支
    release 分支: 发布分支，准备待发布版本的分支，存在多个，版本发布之后删除
    feature 分支: 功能分支，完成特定功能开发的分支，存在多个，功能合并之后删除
    hotfix 分支: 紧急热修复分支，存在多个，紧急版本发布之后删除
    
    
#### git  提交信息规范

    <type>(<scope>): <subject>

    type 类型，提交的类别
    
    - feat: 新功能
    - fix: 修复 bug
    - docs: 文档变动
    - style: 格式调整，对代码实际运行没有改动，例如添加空行、格式化等
    - refactor: bug 修复和添加新功能之外的代码改动
    - perf: 提升性能的改动
    - test: 添加或修正测试代码
    - chore: 构建过程或辅助工具和库（如文档生成）的更改
    
    scope 修改范围
    
    - 主要是这次修改涉及到的部分，简单概括，例如 login、train-order
    
    subject 修改的描述
    
    - 具体的修改描述信息
    
    例子:
    
    feat(detail): 详情页修改样式
    fix(login): 登录页面错误处理
    test(list): 列表页添加测试代码
    
    
    

    