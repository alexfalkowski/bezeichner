# frozen_string_literal: true

When('I request to generate identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::GenerateIdentifiersRequest.new(application: rows['application'], count: rows['count'].to_i)
  @response = Bezeichner::V1.grpc.generate_identifiers(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

When('I request to list applications with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::ListApplicationsRequest.new
  @response = Bezeichner::V1.grpc.list_applications(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

When('I request to map identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::MapIdentifiersRequest.new(application: rows['application'], ids: rows['request'].split(','))
  @response = Bezeichner::V1.grpc.map_identifiers(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

When('I request to map {int} identifiers with gRPC:') do |count, table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::MapIdentifiersRequest.new(application: rows['application'], ids: count.times.map { SecureRandom.hex })
  @response = Bezeichner::V1.grpc.map_identifiers(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

Then('I should receive configured applications from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(applications(@response.generator_applications)).to eq(mapping(rows['generator_applications']))
  expect(@response.mapper_applications.map(&:name)).to eq(identifiers(rows['mapper_applications']))
  expect(@response.generator_kinds).to eq(identifiers(rows['generator_kinds']))
  expect(@response.limits.generate_count).to eq(rows['generate_count'].to_i)
  expect(@response.limits.map_ids).to eq(rows['map_ids'].to_i)
end

Then('I should receive generated identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(@response.ids.length).to eq(rows['count'].to_i)
  expect(@response.ids).to all(satisfy { |id| id.start_with?("#{rows['application']}_") })
end

Then('I should receive mapped identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(mapped_identifier_results(@response.ids)).to eq(mapped_identifiers(rows['results']))
end

Then('I should receive {int} unmapped identifiers from gRPC') do |count|
  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(mapped_identifier_results(@response.ids)).to all(satisfy { |result| !result.key?('mapped') })
  expect(@response.ids.length).to eq(count)
end

Then('I should receive a not found error from gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive an internal error from gRPC') do
  expect(@response).to be_a(GRPC::Internal)
end

Then('I should receive an invalid argument error from gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end
