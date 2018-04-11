package txf

import (
	"fmt"
	"io"
)

type RefNum int

func (r RefNum) Write(w io.Writer) {
	fmt.Fprintf(w, "N%03d\n", int(r))
}

const (
	Form_1040                          RefNum = 256
	F1040_Other_income_misc_           RefNum = 257
	F1040_Sick_pay_or_disab_pay        RefNum = 258
	F1040_Prizes_awards                RefNum = 259
	F1040_Alimony_received             RefNum = 261
	F1040_IRA_contribution_self        RefNum = 262
	F1040_IRA_contribution_spouse      RefNum = 481
	F1040_Med_savings_contrib_self     RefNum = 607
	F1040_Med_savings_contrib_spouse   RefNum = 608
	F1040_SIMPLE_contrib_self          RefNum = 609
	F1040_SIMPLE_contrib_spouse        RefNum = 610
	F1040_Keogh_deduction              RefNum = 263
	F1040_Keogh_deduction_spouse       RefNum = 516
	F1040_SEP_deduction                RefNum = 517
	F1040_SEP_deduction_spouse         RefNum = 518
	F1040_Alimony_paid                 RefNum = 264
	F1040_Early_withd_penalty          RefNum = 265
	F1040_Soc_Sec_income               RefNum = 266
	F1040_Fed_tax_w_h_Soc_Sec_inc_     RefNum = 611
	F1040_Soc_Sec_income_spouse        RefNum = 483
	F1040_Fed_tax_w_h_Soc_Sec_spouse   RefNum = 612
	F1040_RR_retirement_income         RefNum = 519
	F1040_Fed_tax_w_h_RR_retire_inc_   RefNum = 613
	F1040_RR_retirement_inc_spouse     RefNum = 520
	F1040_Fed_tax_w_h_RR_retire_spouse RefNum = 614
	F1040_Fed_estimated_tax            RefNum = 268
	F1040_Fed_est_tax_qrtrly           RefNum = 521
	F1040_Taxable_fringe_benefits      RefNum = 269
	F1040_Student_loan_interest        RefNum = 636
	F1040_Educator_expenses_self       RefNum = 680
	F1040_Educator_expenses_spouse     RefNum = 681

	Schedule_A                   RefNum = 270
	A_Subscriptions              RefNum = 271
	A_Gambling_losses            RefNum = 272
	A_Medicine_and_drugs         RefNum = 273
	A_Med_transport_lodging      RefNum = 274
	A_Doctors_dentists_hosp_     RefNum = 484
	A_State_taxes                RefNum = 275
	A_Local_taxes                RefNum = 544
	A_State_est_tax_qrtrly       RefNum = 522
	A_Real_estate_tax            RefNum = 276
	A_Personal_property_taxes    RefNum = 535
	A_Other_taxes                RefNum = 277
	A_Cash_charity_contributions RefNum = 280
	A_Non_cash_charity_contrib   RefNum = 485
	A_Tax_preparation_fees       RefNum = 281
	A_Investment_man_fees        RefNum = 282
	A_Home_mortgage_int_1098_    RefNum = 283
	A_Home_mortgage_int_no_1098_ RefNum = 545
	A_Points_paid                RefNum = 284
	A_Misc_subject_to_2_         RefNum = 486
	A_Misc_not_subject_to_2_     RefNum = 523

	Schedule_B                    RefNum = 285
	B_Total_dividend_income       RefNum = 286
	B_Qualified_dividend          RefNum = 683
	B_Div_inc_non_taxable         RefNum = 487
	B_Div_inc_capital_gain_dist   RefNum = 488
	B_Fed_tax_w_h_dividend_income RefNum = 615
	B_Interest_income             RefNum = 287
	B_US_govt_interest            RefNum = 288
	B_State_and_mun_bond_int_     RefNum = 289
	B_TE_priv_act_bond_int        RefNum = 290
	B_Interest_inc_non_taxable    RefNum = 489
	B_Int_inc_fed_tax_state_non   RefNum = 490
	B_Int_inc_fed_non_state_tax   RefNum = 491
	B_Int_inc_orig_issue_disc     RefNum = 492
	B_Fed_tax_w_h_interest_income RefNum = 616
	B_Seller_financed_mtge_int_   RefNum = 524

	Schedule_C                  RefNum = 291
	C_Spouse                    RefNum = 292
	C_Gross_receipts            RefNum = 293
	C_Meals_and_entertainment   RefNum = 294
	C_Purchases_cost_of_goods   RefNum = 493
	C_Labor_cost_of_goods_sold  RefNum = 494
	C_Materials_cost_of_goods   RefNum = 495
	C_Other_costs_cost_of_goods RefNum = 496
	C_Returns_and_allowances    RefNum = 296
	C_Wages_paid                RefNum = 297
	C_Legal_and_professional    RefNum = 298
	C_Rent_on_vehicles_mach_eq  RefNum = 299
	C_Rent_on_other_bus_prop    RefNum = 300
	C_Supplies                  RefNum = 301
	C_Other_business_expense    RefNum = 302
	C_Other_business_income     RefNum = 303
	C_Advertising               RefNum = 304
	C_Bad_debts                 RefNum = 305
	C_Car_and_truck_expenses    RefNum = 306
	C_Commissions_and_fees      RefNum = 307
	C_Contract_labor            RefNum = 685
	C_Employee_benefits_progs_  RefNum = 308
	C_Depletion                 RefNum = 309
	C_Insurance_not_health_     RefNum = 310
	C_Interest_expense_mortgage RefNum = 311
	C_Interest_expense_other    RefNum = 312
	C_Office_expense            RefNum = 313
	C_Pension_and_profit_shrg   RefNum = 314
	C_Repairs_and_maintenance   RefNum = 315
	C_Taxes_and_licenses        RefNum = 316
	C_Travel                    RefNum = 317
	C_Utilities                 RefNum = 318
	C_Principal_business_prof   RefNum = 319

	Schedule_C_EZ                RefNum = 497
	C_EZ_Spouse                  RefNum = 498
	C_EZ_Gross_receipts          RefNum = 499
	C_EZ_Total_expenses          RefNum = 500
	C_EZ_Principal_business_prof RefNum = 501

	Schedule_D                 RefNum = 320
	D_ST_gain_loss_8949_Copy_A RefNum = 321
	D_ST_gain_loss_8949_Copy_B RefNum = 711
	D_ST_gain_loss_8949_Copy_C RefNum = 712
	D_LT_gain_loss_8949_Copy_A RefNum = 323
	D_LT_gain_loss_8949_Copy_B RefNum = 713
	D_LT_gain_loss_8949_Copy_C RefNum = 714
	D_Short_Long_8949_Copy_A   RefNum = 673 // Unknown holding
	D_Short_Long_8949_Copy_B   RefNum = 715
	D_Short_Long_8949_Copy_C   RefNum = 716
	D_Wash_Sale_8949_Copy_A    RefNum = 682
	D_Wash_Sale_8949_Copy_B    RefNum = 718
	D_Pct28_cap_gain           RefNum = 644
	D_Unrec_sec_1250           RefNum = 645
	D_Sec_1202_gain            RefNum = 646

	Schedule_E                  RefNum = 325
	E_Rents_received            RefNum = 326
	E_Royalties_received        RefNum = 327
	E_Advertising               RefNum = 328
	E_Auto_and_travel           RefNum = 329
	E_Cleaning_and_maintenance  RefNum = 330
	E_Commissions               RefNum = 331
	E_Insurance                 RefNum = 332
	E_Legal_and_professional    RefNum = 333
	E_Management_fees           RefNum = 502
	E_Mortgage_interest_exp     RefNum = 334
	E_Other_interest_expense    RefNum = 335
	E_Repairs                   RefNum = 336
	E_Supplies                  RefNum = 337
	E_Taxes                     RefNum = 338
	E_Utilities                 RefNum = 339
	E_Other_expenses            RefNum = 341
	E_Kind_location_of_property RefNum = 342

	Schedule_F                     RefNum = 343
	F_Spouse                       RefNum = 514
	F_Car_and_truck_expense        RefNum = 543
	F_Labor_hired                  RefNum = 344
	F_Repairs_and_maintenance      RefNum = 345
	F_Interest_expense_mort        RefNum = 346
	F_Interest_expense_other       RefNum = 347
	F_Rent_land_animals            RefNum = 348
	F_Rent_veh_mach_equip_         RefNum = 349
	F_Feed_purchased               RefNum = 350
	F_Seed_and_plants_purchased    RefNum = 351
	F_Fertilizers_and_lime         RefNum = 352
	F_Supplies_purchased           RefNum = 353
	F_Vet_breeding_and_medicine    RefNum = 355
	F_Gasoline_fuel_oil            RefNum = 356
	F_Storage_and_warehousing      RefNum = 357
	F_Taxes                        RefNum = 358
	F_Insurance_not_health_        RefNum = 359
	F_Utilities                    RefNum = 360
	F_Freight_and_trucking         RefNum = 361
	F_Conservation_expense         RefNum = 362
	F_Pension_and_profit_shrg      RefNum = 363
	F_Employee_benefit_prog_       RefNum = 364
	F_Other_farm_expenses          RefNum = 365
	F_Chemicals                    RefNum = 366
	F_Custom_hire_mach_work_       RefNum = 367
	F_Sales_livestock_crops_raised RefNum = 368
	F_Resales_of_livestock         RefNum = 369
	F_Custom_hire_machine_work_    RefNum = 370
	F_Coop_distributions           RefNum = 371
	F_Agric_pgm_payments           RefNum = 372
	F_CCC_loans_election           RefNum = 373
	F_CCC_loans_forfeited          RefNum = 374
	F_Crop_ins_proceeds_recd       RefNum = 375
	F_Crop_ins_proceeds_defd       RefNum = 376
	F_Other_farm_income            RefNum = 377
	F_Cost_livestock_for_resale    RefNum = 378
	F_Principal_product            RefNum = 379

	Schedule_H             RefNum = 565
	H_Cash_wages_paid      RefNum = 567
	H_Federal_tax_withheld RefNum = 568

	Form_2106                    RefNum = 380
	F2106_Education_expense      RefNum = 381
	F2106_Automobile_exp         RefNum = 382
	F2106_Travel                 RefNum = 383
	F2106_Local_transportation   RefNum = 384
	F2106_Other_bus_exp_         RefNum = 385
	F2106_Meal_and_entertain_exp RefNum = 386
	F2106_Emp_expense_reimb_     RefNum = 387
	F2106_Emp_meal_exp_reimb     RefNum = 388
	F2106_Job_seeking_exp        RefNum = 389
	F2106_Special_clothing_exp   RefNum = 390
	F2106_Emp_home_office_exp    RefNum = 391

	Home_Sale_Worksheets      RefNum = 392
	Selling_price_of_old_home RefNum = 393
	Expense_of_sale           RefNum = 394
	Basis_of_home_sold        RefNum = 395
	Date_old_home_sold        RefNum = 398

	Form_2441                  RefNum = 400
	F2441_Child_care_day_care  RefNum = 401
	F2441_Child_care_household RefNum = 402

	Form_3903                     RefNum = 403
	F3903_Spouse                  RefNum = 526
	F3903_Trans_store_hshld_goods RefNum = 406
	F3903_Travel_and_lodging      RefNum = 407

	Form_4137                       RefNum = 503
	F4137_Spouse                    RefNum = 504
	F4137_Cash_charge_tips_received RefNum = 505

	Form_4684                     RefNum = 412
	F4684_Basis_of_casualty_prop  RefNum = 413
	F4684_Insurance_reimb         RefNum = 414
	F4684_FMV_before_casualty     RefNum = 415
	F4684_FMV_after_casualty      RefNum = 416
	F4684_Description_of_property RefNum = 417

	Form_4797                   RefNum = 418
	F4797_LT_dep_loss_business  RefNum = 419
	F4797_LT_dep_gain_business  RefNum = 420
	F4797_ST_dep_prop_business  RefNum = 421
	F4797_LT_dep_loss_res_rent_ RefNum = 422
	F4797_LT_dep_gain_res_rent_ RefNum = 423
	F4797_ST_dep_prop_res_rent_ RefNum = 424

	Form_4835                          RefNum = 569
	F4835_Spouse                       RefNum = 570
	F4835_Sale_of_livestock_produce    RefNum = 571
	F4835_Total_co_op_distributions    RefNum = 572
	F4835_Agricultural_program_paymnts RefNum = 573
	F4835_CCC_loans_reported_election  RefNum = 574
	F4835_CCC_loans_forfeited_repaid   RefNum = 575
	F4835_Crop_ins_proceeds_rec_d      RefNum = 576
	F4835_Crop_ins_proceeds_deferred   RefNum = 577
	F4835_Other_income                 RefNum = 578
	F4835_Car_and_truck_expenses       RefNum = 579
	F4835_Chemicals                    RefNum = 580
	F4835_Conservation_expenses        RefNum = 581
	F4835_Custom_hire_machine_work_    RefNum = 582
	F4835_Employee_benefit_programs    RefNum = 583
	F4835_Feed_purchased               RefNum = 584
	F4835_Fertilizers_and_lime         RefNum = 585
	F4835_Freight_and_trucking         RefNum = 586
	F4835_Gasoline_fuel_and_oil        RefNum = 587
	F4835_Insurance_other_than_health_ RefNum = 588
	F4835_Interest_expense_mortgage    RefNum = 589
	F4835_Interest_expense_other       RefNum = 590
	F4835_Labor_hired                  RefNum = 591
	F4835_Pension_profit_sharing_plans RefNum = 592
	F4835_Rent_lease_vehicles_equip_   RefNum = 593
	F4835_Rent_lease_land_animals      RefNum = 594
	F4835_Repairs_and_maintenance      RefNum = 595
	F4835_Seeds_and_plants_purchased   RefNum = 596
	F4835_Storage_and_warehousing      RefNum = 597
	F4835_Supplies_purchased           RefNum = 598
	F4835_Taxes                        RefNum = 599
	F4835_Utilities                    RefNum = 600
	F4835_Vet_breeding_medicine        RefNum = 601
	F4835_Other_expenses               RefNum = 602

	Form_4952                         RefNum = 425
	F4952_Investment_interest_expense RefNum = 426

	Form_6252                       RefNum = 427
	F6252_Selling_price             RefNum = 428
	F6252_Debt_assumed_by_buyer     RefNum = 429
	F6252_Basis_of_property_sold    RefNum = 430
	F6252_Depreciation_allowed      RefNum = 431
	F6252_Expenses_of_sale          RefNum = 432
	F6252_Gross_profit_percentage   RefNum = 433
	F6252_Payments_recd_this_year   RefNum = 434
	F6252_Payments_recd_prior_years RefNum = 435
	F6252_Description_of_property   RefNum = 436

	Form_8606                        RefNum = 437
	F8606_Spouse                     RefNum = 515
	F8606_IRAs_value_at_end_of_year  RefNum = 438
	F8606_IRA_contribs_nondeductible RefNum = 439
	F8606_IRA_basis_at_beg_of_year   RefNum = 440

	Form_8815                        RefNum = 441
	F8815_Qual_higher_ed_expenses    RefNum = 442
	F8815_Nontaxable_ed_benefits     RefNum = 443
	F8815_EE_US_svgs_bonds_proceeds  RefNum = 444
	F8815_Post_89_EE_bond_face_value RefNum = 445

	Form_8829                          RefNum = 536
	F8829_Deductible_mortgage_interest RefNum = 537
	F8829_Real_estate_taxes            RefNum = 538
	F8829_Insurance                    RefNum = 539
	F8829_Repairs_and_maintenance      RefNum = 540
	F8829_Utilities                    RefNum = 541
	F8829_Other_expenses               RefNum = 542

	Form_8839                RefNum = 617
	F8839_Adoption_fees      RefNum = 618
	F8839_Court_costs        RefNum = 619
	F8839_Attorney_fees      RefNum = 620
	F8839_Traveling_expenses RefNum = 621
	F8839_Other_expenses     RefNum = 622

	Form_8863                       RefNum = 639
	F8863_Hope_cr_qual_expenses     RefNum = 637
	F8863_Lifetime_cr_qual_expenses RefNum = 638

	Schedule_K_1_Worksheet        RefNum = 446
	K1_Spouse                     RefNum = 447
	K1_Ordinary_income_loss       RefNum = 448
	K1_Rental_real_est_inc_loss   RefNum = 449
	K1_Other_rental_income_loss   RefNum = 450
	K1_Interest_income            RefNum = 451
	K1_Dividends                  RefNum = 452
	K1_Royalties                  RefNum = 527
	K1_Net_ST_capital_gain_loss   RefNum = 453
	K1_Net_LT_capital_gain_loss   RefNum = 454
	K1_Pct28_rate_gain_loss_      RefNum = 674
	K1_Guaranteed_payments        RefNum = 455
	K1_Net_sec_1231_gain_loss     RefNum = 456
	K1_Other_Income_loss_         RefNum = 676
	K1_Tax_exempt_interest        RefNum = 528
	K1_Total_Foreign_Taxes        RefNum = 679
	K1_Partnership_or_S_corp_name RefNum = 457

	W_2                            RefNum = 458
	W2_Salary                      RefNum = 460
	W2_Federal_withholding         RefNum = 461
	W2_Soc_Sec_tax_withholding     RefNum = 462
	W2_Local_withholding           RefNum = 463
	W2_Medicare_tax_withholding    RefNum = 480
	W2_State_withholding           RefNum = 464
	W2_Dependent_care_benefits     RefNum = 465
	W2_Moving_exp_reimb_           RefNum = 267
	W2_Payer                       RefNum = 466
	W2_Salary_spouse               RefNum = 506
	W2_Federal_withholding_spouse  RefNum = 507
	W2_Soc_Sec_tax_withhld_spouse  RefNum = 508
	W2_Local_withholding_spouse    RefNum = 509
	W2_Medicare_tax_withhld_spouse RefNum = 510
	W2_State_withholding_spouse    RefNum = 511
	W2_Dependent_care_ben_spouse   RefNum = 512
	W2_Moving_exp_reimb_spouse     RefNum = 546
	W2_Payer_spouse                RefNum = 513

	W_2G                     RefNum = 547
	W2G_Spouse               RefNum = 548
	W2G_Gross_winnings       RefNum = 549
	W2G_Federal_tax_withheld RefNum = 550
	W2G_State_tax_withheld   RefNum = 551
	W2G_Payer                RefNum = 552

	F1099_MISC                         RefNum = 553
	F1099MISC_Spouse                   RefNum = 554
	F1099MISC_Rents                    RefNum = 555
	F1099MISC_Royalties                RefNum = 556
	F1099MISC_Other_income             RefNum = 557
	F1099MISC_Federal_tax_withheld     RefNum = 558
	F1099MISC_Fishing_boat_proceeds    RefNum = 559
	F1099MISC_Medical_health_payments  RefNum = 560
	F1099MISC_Nonemployee_compensation RefNum = 561
	F1099MISC_Crop_insurance_proceeds  RefNum = 562
	F1099MISC_State_tax_withheld       RefNum = 563
	F1099MISC_Payer                    RefNum = 564
	F1099MISC_State_ID                 RefNum = 654
	F1099MISC_Substitute_payments_     RefNum = 674
	F1099MISC_Gross_proceeds_to_atty_  RefNum = 675
	F1099MISC_Excess_golden_parachute  RefNum = 676

	F1099R                              RefNum = 473
	F1099R_Spouse                       RefNum = 474
	F1099R_Pension_total_dist_gross     RefNum = 475
	F1099R_Pension_total_dist_taxable   RefNum = 476
	F1099R_Pension_federal_withholding  RefNum = 529
	F1099R_Pension_state_withholding    RefNum = 530
	F1099R_Pension_local_withholding    RefNum = 531
	F1099R_Pension_Distrib_code_7A      RefNum = 667
	F1099R_Pension_Distrib_code_7b      RefNum = 668
	F1099R_Pension_State_ID             RefNum = 669
	F1099R_Pension_Tax_not_determined   RefNum = 670
	F1099R_Pension_Total_distribution   RefNum = 671
	F1099R_Pension_Payer                RefNum = 603
	F1099R_IRA_total_dist_gross         RefNum = 477
	F1099R_IRA_total_dist_taxable       RefNum = 478
	F1099R_IRA_federal_withholding      RefNum = 532
	F1099R_IRA_state_withholding        RefNum = 533
	F1099R_IRA_local_withholding        RefNum = 534
	F1099R_IRA_Distribution_code_7A     RefNum = 655
	F1099R_IRA_Distribution_code_7b     RefNum = 656
	F1099R_IRA_State_ID                 RefNum = 664
	F1099R_IRA_Taxable_not_determined   RefNum = 665
	F1099R_IRA_Total_distribution       RefNum = 666
	F1099R_IRA_Payer                    RefNum = 604
	F1099R_SIMPLE_total_distrib_gross   RefNum = 623
	F1099R_SIMPLE_total_distrib_taxable RefNum = 624
	F1099R_SIMPLE_federal_withholding   RefNum = 625
	F1099R_SIMPLE_state_withholding     RefNum = 626
	F1099R_SIMPLE_local_withholding     RefNum = 627
	F1099R_SIMPLE_Payer                 RefNum = 628
	F1099R_Distribution_code_7A         RefNum = 655
	F1099R_Distribution_code_7b         RefNum = 656
	F1099R_State_ID                     RefNum = 664
	F1099R_Taxable_not_determined       RefNum = 665
	F1099R_Total_distribution           RefNum = 666

	F1099_SA                        RefNum = 629
	F1099SA_Spouse                  RefNum = 630
	F1099SA_gross_distribution      RefNum = 631
	F1099SA_earnings_excess_contrib RefNum = 632
	F1099SA_Payer                   RefNum = 633

	F1099_G                            RefNum = 634
	F1099G_Spouse                      RefNum = 635
	F1099G_Unemployment_comp           RefNum = 479
	F1099G_Unemployment_comp_repaid    RefNum = 605
	F1099G_State_and_local_refunds     RefNum = 260
	F1099G_Fed_tax_w_h_unemploymt_comp RefNum = 606

	F1099_INT                       RefNum = 640
	F1099INT_Foreign_tax_int        RefNum = 641
	F1099INT_Foreign_country_int    RefNum = 642
	F1099INT_Intestment_expense_int RefNum = 653

	F1099_Div                         RefNum = 643
	F1099DIV_NonTaxable_distributions RefNum = 647
	F1099DIV_Investment_Expense_div   RefNum = 648
	F1099DIV_Foreign_tax_div          RefNum = 649
	F1099DIV_Foreign_country_div      RefNum = 650
	F1099DIV_Cash_liquidation         RefNum = 651
	F1099DIV_Non_cash_liquidation     RefNum = 652

	F1099_OID                       RefNum = 657
	F1099OID_Other_periodic_int_OID RefNum = 658
	F1099OID_Early_wdrawal_pen_OID  RefNum = 659
	F1099OID_Fed_tax_wheld_OID      RefNum = 660
	F1099OID_Description            RefNum = 661
	F1099OID_OID_US_treas_obligl    RefNum = 662
	F1099OID_Investment_Expense_OID RefNum = 663

	F1099_Q                   RefNum = 678
	F1099Q_Gross_distribution RefNum = 672
)
