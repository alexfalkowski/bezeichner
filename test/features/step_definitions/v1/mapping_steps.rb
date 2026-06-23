# frozen_string_literal: true

def mapping(value)
  return {} if value.nil? || value.empty?

  value.split(',').to_h { |entry| entry.split(':', 2) }
end

def identifiers(value)
  return [] if value.nil? || value.empty?

  value.split(',')
end
