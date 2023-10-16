def isIsomorphic(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    t_map = {}
    s_map = {}
