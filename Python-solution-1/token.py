class Side:
    def __init__(self, colour, part):
        self.c = colour
        self.p = part


class Token:
    def __init__(self, north: Side, east: Side, south: Side, west: Side):
        self.north = north
        self.east = east
        self.south = south
        self.west = west
        self.faces = [self.north, self.east, self.south, self.west]

    def rotate(self, rotation: int):
        self.actual_north = self.faces[(0 + rotation) % len(self.faces)]
        self.actual_east = self.faces[(1 + rotation) % len(self.faces)]
        self.actual_south = self.faces[(2 + rotation) % len(self.faces)]
        self.actual_west = self.faces[(3 + rotation) % len(self.faces)]


class Board:
    def __init__(self, tokens: list, rotations: list):
        """Note that the tokens should be given as a list, in order of their placement on the board.
        The same is true of the rotations, which should also be given as an ordered list."""
        self.spot_1 = tokens[0]
        self.spot_2 = tokens[1]
        self.spot_3 = tokens[2]
        self.spot_4 = tokens[3]
        self.spot_5 = tokens[4]
        self.spot_6 = tokens[5]
        self.spot_7 = tokens[6]
        self.spot_8 = tokens[7]
        self.spot_9 = tokens[8]

    def rotate_next(self, spot_to_rotate, rotation):
        pass
