    PROID=4
    # Your script logic here

    echo "Running script..."
    #yes | ./build/terrad tx gov submit-proposal draft_proposal.json --from test0 --keyring-backend test --home mytestnet --fees 5665000uluna 
    
    sleep 2
    
    echo "deposit"
    yes | ./build/terrad tx gov deposit $PROID 20000000uluna --from test0 --keyring-backend test --home mytestnet --fees 5665000uluna  
    
    sleep 2
    
    echo "vote1"
    yes | ./build/terrad tx gov vote $PROID yes --from test0 --keyring-backend test --home mytestnet --fees 5665000uluna
    
    sleep 2

    echo "vote2"
   yes | ./build/terrad tx gov vote $PROID yes --from test1 --keyring-backend test --home mytestnet --fees 5665000uluna 
    
    sleep 2
    
    echo "vote3"
    yes | ./build/terrad tx gov vote $PROID yes --from test2 --keyring-backend test --home mytestnet --fees 5665000uluna

    echo "done"
    ./build/terrad q gov proposal $PROID