def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    s_map, t_map = {}, {}
    for c1, c2 in zip(s, t):
        if(c1 in )


s = "egg"
t = "add"
print(isIsomorphic(s, t))
