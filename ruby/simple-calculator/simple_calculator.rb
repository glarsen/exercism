class SimpleCalculator
  class UnsupportedOperation < StandardError
  end

  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  def self.calculate(first_operand, second_operand, operation)
    first_operand.is_a?(Numeric) && second_operand.is_a?(Numeric) or raise ArgumentError
    raise UnsupportedOperation unless ALLOWED_OPERATIONS.include? operation

    begin
      "#{first_operand} #{operation} #{second_operand} = #{first_operand.send(operation, second_operand)}"
    rescue ZeroDivisionError
      "Division by zero is not allowed."
    end
  end
end
