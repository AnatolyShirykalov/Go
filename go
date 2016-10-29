current_path=$(pwd)
  if [ -z "$1" ]; then
    echo Usage:
    echo
    echo go ac from_path to_path
  else
    from_path=$1
    to_path=$2

    echo cd $from_path
    cd $from_path
    
    echo getting list of commits
    commits=($(git log | grep commit | awk '{print $2}' | tac))
    messages=($(git log | grep -E '^  ' | tac))
    for commit in $commits; do
      echo "  $commit"
    done
    echo done
    echo
        
    echo rm -rf $to_path
    rm -rf $to_path
    echo mkdir $to_path
    mkdir $to_path
    echo cd $to_path
    cd $to_path
    echo git init
    git init

    for (( i=1; i<=${#commits} ; i++ )) do
      commit=${commits[i]}
      message=${messages[i]}
      echo cd $from_path
      cd $from_path
      echo git checkout $commit
      git checkout $commit
      echo rsync -az --exclude '.git' $from_path $to_path
      rsync -az --exclude '.git' $from_path $to_path
      echo cd $to_path
      cd $to_path
      echo git add .
      git add .
      echo git commit -m $message
      git commit -m $message
      sleep $((900+$RANDOM%100))
    done

    echo cd $current_path
    cd $current_path

  fi
