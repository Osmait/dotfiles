def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    s_map, t_map = {}, {}
    for c1, c2 in zip(s, t):
        if ((c2 in t_map and t_map[c2] != c1) or (c1 in s_map and s_map[c1] != c2)):
            return False
        s_map[c2] = c1
        t_map[c1] = c2
    return True


s = "egg"
t = "add"
print(isIsomorphic(s, t))
