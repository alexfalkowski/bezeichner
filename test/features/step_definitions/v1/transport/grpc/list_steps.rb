# frozen_string_literal: true

When('I request to list applications with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::ListApplicationsRequest.new
  @response = Bezeichner::V1.grpc.list_applications(request, Bezeichner.grpc_options(metadata:))
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
