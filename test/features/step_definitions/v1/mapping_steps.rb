# frozen_string_literal: true

def mapping(value)
  return {} if value.nil? || value.empty?

  value.split(',').to_h { |entry| entry.split(':', 2) }
end

def identifiers(value)
  return [] if value.nil? || value.empty?

  value.split(',')
end

def applications(value)
  value.to_h do |app|
    name = app.respond_to?(:name) ? app.name : app['name']
    kind = app.respond_to?(:kind) ? app.kind : app['kind']

    [name, kind]
  end
end
