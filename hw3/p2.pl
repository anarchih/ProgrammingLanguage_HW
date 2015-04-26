

isConnected(A,B) :-   
  move(A,B,[]) 
  ,!.           

move(A,B,V) :-       
  edge(A,X) ,        
  not(member(X,V)) , 
  (                  
    B = X            
  ;                  
    move(X,B,[A|V])  
  )                  
  ,!.                 

inputR(N) :-
    (N > 0) ->
        readln(X),
        nth0(0, X, X1),
        nth0(1, X, X2),
        asserta((edge(X1, X2))), 
        inputR(N - 1);
        !.

inputQ(N) :-
    (N > 0) ->
        readln(X),
        nth0(0, X, X1),
        nth0(1, X, X2),
        ((isConnected(X1, X2)) ->
            write('Yes'),
            nl,
            inputQ(N - 1),
            !;
            
            write('No'),
            nl,
            inputQ(N - 1),
            !);

        !.

main :-
    readln(X),
    nth0(1, X, X1),
    inputR(X1),
    readln(Y),
    inputQ(Y),
    halt.

:- initialization(main).
