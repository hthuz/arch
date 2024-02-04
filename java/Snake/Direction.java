public enum Direction {
    UP() {
        @Override
        public boolean isUP() {return true;}
    },
    DOWN() {
        @Override
        public boolean isDOWN() {return true;}
    },
    LEFT() {
        @Override
        public boolean isLEFT() {return true;}
    },

    RIGHT() {
        @Override
        public boolean isRIGHT() {return true;}
    };
    Direction() {
    }
    public boolean isUP() {return false;}
    public boolean isDOWN() {return false;}
    public boolean isLEFT() {return false;}
    public boolean isRIGHT() {return false;}

}
