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

Then('I should receive generated identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(@response.ids.length).to eq(rows['count'].to_i)
  expect(@response.ids).to all(satisfy { |id| id.start_with?("#{rows['application']}_") })
end
