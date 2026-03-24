# frozen_string_literal: true

When('I request to generate identifiers with gRPC which performs in {int} ms') do |time|
  metadata = { 'request-id' => SecureRandom.uuid }
  request = Bezeichner::V1::GenerateIdentifiersRequest.new(application: 'ulid', count: 2)

  expect { Bezeichner::V1.grpc.generate_identifiers(request, { metadata: }) }.to perform_under(time).ms
end
