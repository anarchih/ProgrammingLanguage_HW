edge(1, 2).
edge(2, 1).
edge(2, 3).
edge(3, 2).
edge(1, 4).
edge(4, 1).
edge(4, 5).
edge(5, 4).
edge(4, 6).
edge(6, 4).

isConnected(X,X,N):-!.
isConnected(X,Y,N):-(N>0)->edge(X,Y),write('Yes')!;write('No'),nl,halt.
isConnected(X, Z,N) :- (N>0)->edge(X, Y),write(Y), isConnected(Y, Z, N -1),!;write('No'),nl,halt.
test :-
    isConnected(1, 6, 6),
    write('Yes'),
    nl,
    halt.

