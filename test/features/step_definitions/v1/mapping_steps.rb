# frozen_string_literal: true

def mapping(value)
  return {} if value.nil? || value.empty?

  value.split(',').to_h { |entry| entry.split(':', 2) }
end

def identifiers(value)
  return [] if value.nil? || value.empty?

  value.split(',')
end

def mapped_identifiers(value)
  return [] if value.nil? || value.empty?

  value.split(',').map do |entry|
    id, mapped = entry.split(':', 2)
    result = { 'id' => id }
    result['mapped'] = mapped unless mapped.nil? || mapped.empty?
    result
  end
end

def mapped_identifier_results(value)
  value.map do |id|
    result = { 'id' => field(id, 'id') }
    mapped = field(id, 'mapped')
    result['mapped'] = mapped if field?(id, 'mapped') && !mapped.empty?
    result
  end
end

def applications(value)
  value.to_h do |app|
    name = app.respond_to?(:name) ? app.name : app['name']
    kind = app.respond_to?(:kind) ? app.kind : app['kind']

    [name, kind]
  end
end

def field(value, name)
  value.respond_to?(name) ? value.public_send(name) : value[name]
end

def field?(value, name)
  predicate = "has_#{name}?"
  return value.public_send(predicate) if value.respond_to?(predicate)

  value.key?(name)
end
