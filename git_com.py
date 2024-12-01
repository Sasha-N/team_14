import os
def git_add(file_local="."):
#  !git add $file_local
  os.system(f"git add {file_local}")
def git_clone(repo_url, repo_up):
#  %cd repo_up
#  !git clone $repo_url #{repo_url}
  os.chdir(repo_up)
  os.system(f"git clone {repo_url}")
def git_clone_private(user_name, repo_name_git, repo_up, key):
  #git clone https://<PAT>@github.com/<username>/<repo>.git
  os.chdir(repo_up)
  repo_url=f"https://{user_name}:{key}@github.com/{user_name}/{repo_name_git}"
  os.system(f"git clone {repo_url}")
def git_commit(message):
#  !git commit -m '1 commit message'
  os.system(f"git commit -m '{message}'")  
def git_config(user_name):
#  !git config --global user.email '$uname@gmail.com'
#  !git config --global user.name '$uname'
  os.system(f"git config --global user.email {user_name}@gmail.com")
  os.system(f"git config --global user.name {user_name}")
def git_print():
  print("Тест пройден - модуль подключен")
def git_pull(branch, remote_branch):
#  !git pull $branch main
  os.system(f"git pull {branch} {remote_branch}")
def git_push(branch, remote_branch):
#  !git push $branch main
  os.system(f"git push {branch} {remote_branch}")
def git_remote(myRepo, repo_name, branch, user_name, public_token):
  os.chdir(myRepo)
  repo_name_git = repo_name+".git"
  repo_url=f"https://{user_name}:{public_token}@github.com/{user_name}/{repo_name_git}"
#  branch = "origin511"
#  !git remote add $branch https://$uname:$public_token@github.com/$uname/$repo_name_git
  os.system(f"git remote add {branch} {repo_url}")
def git_remove(branch):
#  !git remote remove origin #main #origin
#  !git remote -v
  os.system(f"git remote remove {branch}")
  os.system("git remote -v")
