def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    t_map = {}
    s_map = {}

    for i in range(0, len(s)-1):
        s_map[s[i]] = t[i]
        t_map[t[i]] = s[i]
    if len(t_map) == len(t) and len(s_map) == len(s):
        return True
    return False


s = "egg"
t = "add"
print(isIsomorphic(s, t))
