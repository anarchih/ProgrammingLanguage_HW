isEqual(X, Y):-
    X == Y.
p(X, Y, Z) :-
    (isEqual(X, Y) ; parent(X, Y)) ->
        Z = X;
        parent(T, X),
        p(T, Y, Z),!.


inputQ(N) :-
    (N > 0) ->
        readln(X),
        nth0(0, X, X1),
        nth0(1, X, X2),
        p(X1, X2, Z),
        write(Z),
        nl,
        inputQ(N - 1);
        !.


inputPar(X) :-
    (nth0(0, X, X1), nth0(1, X, X2)) ->
        asserta((parent(X1, X2))),
        aaa(T);
        nth0(0, X, X1),
        inputQ(X1),
        !.

aaa(X) :-
    readln(X),
    inputPar(X).
main :-
    readln(W),
    aaa(X),
    halt.

:- initialization(main).
