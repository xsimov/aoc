class Coords
  attr_accessor :x, :y, :facing, :history

  def initialize
    @x = 0
    @y = 0
    @facing = Facing.new(:N)
    @history = [[0, 0]]
  end

  def execute(command)
    case command
    when /L/
      @facing.turn_left
    when /R/
      @facing.turn_right
    end
    number = command.match(/\d+/)[0].to_i
    number.times do
      send("move_#{facing.current_pole.downcase}")
      record_current_position
    end
  end

  def move_n
    @y += 1
  end

  def move_e
    @x += 1
  end

  def move_w
    @x -= 1
  end

  def move_s
    @y -= 1
  end

  def current_position
    [@x, @y]
  end

  def record_current_position
    @first_duplicated = current_position if !@first_duplicated && @history.include?(current_position)
    @history << current_position
  end

  def first_duplicated_visit
    @first_duplicated
  end
end


class Facing
  POLES = [:N, :E, :S, :W]

  def initialize(initial_pole)
    @current = POLES.index(initial_pole)
  end

  def turn_left
    @current = (@current - 1) % 4
    current_pole
  end

  def turn_right
    @current = (@current + 1) % 4
    current_pole
  end

  def current_pole
    POLES[@current]
  end
end


COMMANDS = "L2, L3, L3, L4, R1, R2, L3, R3, R3, L1, L3, R2, R3, L3, R4, R3, R3, L1, L4, R4, L2, R5, R1, L5, R1, R3, L5, R2, L2, R2, R1, L1, L3, L3, R4, R5, R4, L1, L189, L2, R2, L5, R5, R45, L3, R4, R77, L1, R1, R194, R2, L5, L3, L2, L1, R5, L3, L3, L5, L5, L5, R2, L1, L2, L3, R2, R5, R4, L2, R3, R5, L2, L2, R3, L3, L2, L1, L3, R5, R4, R3, R2, L1, R2, L5, R4, L5, L4, R4, L2, R5, L3, L2, R4, L1, L2, R2, R3, L2, L5, R1, R1, R3, R4, R1, R2, R4, R5, L3, L5, L3, L3, R5, R4, R1, L3, R1, L3, R3, R3, R3, L1, R3, R4, L5, L3, L1, L5, L4, R4, R1, L4, R3, R3, R5, R4, R3, R3, L1, L2, R1, L4, L4, L3, L4, L3, L5, R2, R4, L2".split(', ')

