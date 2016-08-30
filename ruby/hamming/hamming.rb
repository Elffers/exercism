class Hamming
  VERSION = 2

  def self.compute left, right
    raise ArgumentError.new('Strings not equal length') unless left.size == right.size

    left.split('').zip(right.split('')).count{|l,r| l != r}
  end

end
