<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('Two Columns Multiple'),
  'category' => t('Bootstrap'),
  'icon' => 'bootstrap_twocol_multiple.png',
  'theme' => 'bootstrap_twocol_multiple',
  'admin css' => 'bootstrap_twocol_multiple.css',
  'regions' => array(
    'top' => t('Top'),
    'left1' => t('Left1'),
    'right1' => t('Right1'),
    'left2' => t('Left2'),
    'right2' => t('Right2'),
    'left3' => t('Left3'),
    'right3' => t('Right3'),
    'bottom' => t('Bottom'),
  ),
);
